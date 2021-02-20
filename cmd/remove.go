package cmd

import (
	"fmt"
	"strconv"

	"github.com/grpit/todo-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// removeCmd represents the delete command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo from the list to todos.",
	Long:  `Remove a todo from the list to todos.`,
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

		utils.ReadJSONFromFile(&projectData, path)

		todoLength := len(projectData.Todos)
		if todoLength < todoID {
			fmt.Println(Colors["Red"] + "Todo with given ID does not exist" + Colors["Reset"])
			return nil
		}

		todoID = todoID - 1
		todos := projectData.Todos
		todoName := projectData.Todos[todoID].Name

		projectData.Todos = append(todos[:todoID], todos[todoID+1:]...)

		err = utils.WriteJSONToFile(projectData, path)

		if err != nil {
			return err
		}
		fmt.Printf("%sDeleted todo %s%s%s.%s\n", Colors["Red"], Colors["Cyan"], todoName, Colors["Red"], Colors["Reset"])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
