package cmd

import (
	"fmt"

	"github.com/grpit/todo-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show a list of added todos.",
	Long:  `Show a list of added todos.`,
	Run: func(cmd *cobra.Command, args []string) {

		_, path := getFilePath()
		projectData := Project{}

		utils.ReadJSONFromFile(&projectData, path)

		if len(projectData.Todos) == 0 {
			fmt.Println(Colors["Yellow"] + "No Todos Yet." + Colors["Reset"])
		}

		for _, todo := range projectData.Todos {
			fmt.Println(Colors["Green"] + todo.Name + Colors["Reset"])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
