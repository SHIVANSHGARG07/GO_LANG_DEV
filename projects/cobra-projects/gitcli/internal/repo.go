package internal

import "os"
import "fmt"
import "path/filepath"

const RepoDir = ".gitcli"

// if found then error
func IsRepo() bool {
	_, err := os.Stat(RepoDir)
	return err == nil
}

func InitRepo() error {
	// already repo
	if IsRepo() {
		return fmt.Errorf("Repository already exists")
	}

	// create .gitcli dir
	if err := os.Mkdir(RepoDir, 0755); err != nil {
		return err
	}

	// create empty files
	files := []string{"commits.json", "config.json", "remotes.json"}

	// write each file using os
	for _, file := range files {
		path := filepath.Join(RepoDir, file)
		if err := os.WriteFile(path, []byte("{}"), 0644); err != nil {
			return err
		}
	}
	return nil
}

// get repo path

func GetRepoPath() (string, error) {

	if !IsRepo() {
		return "", fmt.Errorf("Not a repository")
	}

	return RepoDir, nil
}
