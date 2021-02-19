package vcs

import (
	"log"
	"strings"
	"testing"
)

func TestRepoFound(t *testing.T) {
	path := GetRepo()
	expectedPath := "todo-cli"

	if !strings.Contains(path, expectedPath) {
		log.Fatalln("Error: Root folder not the same.")
	}

}
