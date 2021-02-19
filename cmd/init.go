package cmd

import (
	"bufio"
	"fmt"
	"os"
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

var filename = ".todo.json"

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiates a project in current folder or git root.",
	Long:  `Initiates a project in current folder or git root.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, path := getFilePath()

		if err := utils.CreateFileWithPerm(path); err != nil {
			fmt.Println(Colors["Red"] + "A project has already been initialised." + Colors["Reset"])
			return nil
		}

		fmt.Print(Colors["Blue"] + "What would you like to name your project ? " + Colors["Reset"])

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		projectName := scanner.Text()

		projectJSON := Project{
			Name:    projectName,
			Todos:   []Todo{},
			Created: time.Now().UTC().String(),
		}

		if err := utils.WriteJSONToFile(projectJSON, path); err != nil {
			return err
		}

		vcs.AddIfNotPresent(filepath.Join(repo, ".gitignore"), filename)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func getFilePath() (string, string) {
	repo := vcs.GetRepo()
	path := filepath.Join(repo, filename)
	return repo, path
}
