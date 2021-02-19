package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/grpit/todo-cli/pkg/utils"
	"github.com/grpit/todo-cli/pkg/vcs"
	"github.com/spf13/cobra"
)

// Todo : Type to keep todos
type Todo struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// Project : Type to keep project data
type Project struct {
	Name    string `json:"name"`
	Todos   []Todo `json:"todos"`
	Created string `json:"createdAt"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiates a project in current folder or git root.",
	Long:  `Initiates a project in current folder or git root.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var filename = ".todo.json"
		var project string
		repo := vcs.GetRepo()

		fmt.Print(Colors["Blue"] + "What would you like to name your project ? " + Colors["Reset"])
		fmt.Scanln(&project)

		projectJSON := Project{
			Name:    project,
			Todos:   []Todo{},
			Created: time.Now().UTC().String(),
		}

		if err := utils.WriteJSONToFile(projectJSON, filepath.Join(repo, filename)); err != nil {
			return err
		}

		vcs.AddIfNotPresent(filepath.Join(repo, ".gitignore"), filename)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
