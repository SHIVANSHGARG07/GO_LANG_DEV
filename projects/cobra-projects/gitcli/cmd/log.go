// shows commit history

package cmd

import (
	"fmt"

	"github.com/SHIVANSHGARG07/gitcli/internal"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{

	Use:   "log",
	Short: "Show commit history",
	Long:  "Show commit history",

	RunE: func(cmd *cobra.Command, args []string) error {

		// check if repo is there done by newcommit store
		store, err := internal.NewCommitStore()
		if err != nil {
			return err
		}

		// load data from file/internal
		if err := store.Load(); err != nil {
			return err
		}

		// check if not null
		if len(store.Commits) == 0 {
			fmt.Println("No commits found")
			return nil
		}

		// print data in reverse loop
		for i := len(store.Commits) - 1; i >= 0; i-- {
			commit := store.Commits[i]

			// display each commit
			fmt.Printf("commit %s\n", commit.Hash)
			fmt.Printf("Author: %s\n", commit.Author)
			fmt.Printf("Date: %s\n", commit.Date)
			fmt.Printf("    %s\n", commit.Message)
			fmt.Println()

		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
