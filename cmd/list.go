package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show a list of added todos.",
	Long:  `Show a list of added todos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\033[31mlist called\033[0m")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
