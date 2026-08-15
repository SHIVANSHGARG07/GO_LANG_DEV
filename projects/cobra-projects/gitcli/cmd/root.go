package cmd

import (
	"fmt"
	"os"

	"github.com/SHIVANSHGARG07/gitcli/cmd/remote_cmd"
	"github.com/spf13/cobra"
)

var (
	verbose  bool
	repoPath string
)

var rootCmd = &cobra.Command{
	Use:   "gitcli",
	Short: "A simple Git CLI",
	Long:  `A simple Git CLI to demonstrate nested commands`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(remote_cmd.RemoteCmd)

	// persistent flags (for all subcommands)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.PersistentFlags().StringVarP(&repoPath, "repo", "r", "", "Path to repository")
}
