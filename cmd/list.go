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
			if todo.Done {
				fmt.Printf(" [%sx%s] %s%s%s\n", Colors["Red"], Colors["Reset"], Colors["Red"], todo.Name, Colors["Reset"])
			} else {
				fmt.Printf(" [ ] %s%s%s\n", Colors["Cyan"], todo.Name, Colors["Reset"])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
