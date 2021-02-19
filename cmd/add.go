package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the create command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a todo to the current project.",
	Long:  `Adds a todo to the current project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
