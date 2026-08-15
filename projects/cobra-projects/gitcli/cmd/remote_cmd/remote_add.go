package remote_cmd

import (
	"fmt"

	"github.com/SHIVANSHGARG07/gitcli/internal"
	"github.com/spf13/cobra"
)

var remoteAddCmd = &cobra.Command{
	Use:   "add [name] [url]",
	Short: "Add a remote repository",
	Long:  "Add a remote repository with the given name and URL",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		name := args[0]
		url := args[1]

		// Create store (validates repo exists)
		store, err := internal.NewRemoteStore()
		if err != nil {
			return err
		}

		// Load existing remotes
		if err := store.Load(); err != nil {
			return err
		}

		// Add new remote (in memory)
		if err := store.Add(name, url); err != nil {
			return err
		}

		// Save to file
		if err := store.Save(); err != nil {
			return err
		}

		fmt.Printf("Remote '%s' added -> %s\n", name, url)
		return nil
	},
}

func init() {
	RemoteCmd.AddCommand(remoteAddCmd)
}
