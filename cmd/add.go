package cmd

import (
	"fmt"

	"github.com/grpit/todo-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the create command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a todo to the current project.",
	Long:  `Adds a todo to the current project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println(Colors["Red"] + "Please pass a name for the todo in quotes." + Colors["Reset"])
		}

		todoName := args[0]
		_, path := getFilePath()
		projectData := Project{}

		utils.ReadJSONFromFile(&projectData, path)

		todo := Todo{
			Name: todoName,
			Done: false,
		}

		todos := append(projectData.Todos, todo)
		projectData.Todos = todos
		err := utils.WriteJSONToFile(projectData, path)

		if err != nil {
			return err
		}
		fmt.Printf("%sAdded %s%s%s Successfully.%s\n", Colors["Green"], Colors["Cyan"], todo.Name, Colors["Green"], Colors["Reset"])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
