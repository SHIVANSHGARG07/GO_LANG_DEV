package cmd

import (
	"fmt"
	"github.com/SHIVANSHGARG07/gitcli/internal"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a repository",
	Long:  "Initialize a new repository in the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := internal.InitRepo(); err != nil {
			return err
		}
		fmt.Println("Repository initialized")

		// access persistent flags
		if verbose {
			fmt.Println("[verbose] Created commits.json, config.json, remotes.json")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
