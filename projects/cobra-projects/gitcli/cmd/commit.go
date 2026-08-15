package cmd

import (
	"fmt"

	"github.com/SHIVANSHGARG07/gitcli/internal"
	"github.com/spf13/cobra"
)

var (
	commitMessage string
	commitAuthor  string
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a new commit",
	Long:  "Create a new commit with the given message and author",
	RunE: func(cmd *cobra.Command, args []string) error {

		// validate input
		if commitMessage == "" {
			return fmt.Errorf("commit message is required")
		}

		// create store (auto validates repo exists or not)
		store, err := internal.NewCommitStore()
		if err != nil {
			return err
		}

		// load existing commits
		if err := store.Load(); err != nil {
			return err
		}

		if verbose {
			fmt.Printf("[verbose] Loaded %d existing commits\n", len(store.Commits))
		}

		// add new commit
		hash := store.AddCommit(commitMessage, commitAuthor)

		// save to file
		if err := store.Save(); err != nil {
			return err
		}

		// show success message
		fmt.Printf("[%s] %s\n", hash, commitMessage)

		if verbose {
			fmt.Printf("[verbose] Commit saved to .gitcli/commits.json\n")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// define flags
	commitCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "Commit message (required)")
	commitCmd.Flags().StringVarP(&commitAuthor, "author", "a", "user", "Commit author")

	// mark message as required
	commitCmd.MarkFlagRequired("message")
}
