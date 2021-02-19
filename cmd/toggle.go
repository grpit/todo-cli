package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggles the current status of a todo.",
	Long:  `Toggles the current status of a todo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("toggle called")
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
