package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the delete command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo from the list to todos.",
	Long:  `Remove a todo from the list to todos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
