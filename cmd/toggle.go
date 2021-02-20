package cmd

import (
	"fmt"
	"strconv"

	"github.com/grpit/todo-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggles the current status of a todo.",
	Long:  `Toggles the current status of a todo.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println(Colors["Red"] + "Please pass the ID for the todo." + Colors["Reset"])
		}

		todoID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(Colors["Red"] + "Enter a valid todo ID to delete." + Colors["Reset"])
			return nil
		}

		_, path := getFilePath()
		projectData := Project{}

		err = utils.ReadJSONFromFile(&projectData, path)

		if err != nil {
			return err
		}

		todoLength := len(projectData.Todos)
		if todoLength < todoID {
			fmt.Println(Colors["Red"] + "Todo with given ID does not exist" + Colors["Reset"])
			return nil
		}

		todoID = todoID - 1
		todo := projectData.Todos[todoID]
		projectData.Todos[todoID].Done = !projectData.Todos[todoID].Done
		err = utils.WriteJSONToFile(projectData, path)

		if err != nil {
			return err
		}
		if todo.Done {
			fmt.Printf("%sSuccessfully marked todo %s%s%s as Not Done.%s\n", Colors["Green"], Colors["Cyan"], todo.Name, Colors["Green"], Colors["Reset"])
		} else {
			fmt.Printf("%sSuccessfully marked todo %s%s%s as Done.%s\n", Colors["Green"], Colors["Cyan"], todo.Name, Colors["Green"], Colors["Reset"])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
