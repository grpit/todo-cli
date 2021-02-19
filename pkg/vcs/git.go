package vcs

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

// GetRepo : Executes a git command to identify the top level dir.
func GetRepo() string {
	root, err := findRoot()

	if err != nil {
		return "./"
	}

	return root
}

// AddIfNotPresent : Adds todos file to gitignore
func AddIfNotPresent(path string, value string) {
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	isPresent := false

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == value {
			isPresent = true
			break
		}
	}

	if !isPresent {
		file.WriteString(value + "\n")
	}

}

// Gets the root path for the current dir.
func findRoot() (string, error) {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(path)), nil
}
