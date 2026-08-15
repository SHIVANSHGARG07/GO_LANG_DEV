package remote_cmd

import (
	"fmt"

	"github.com/SHIVANSHGARG07/gitcli/internal"
	"github.com/spf13/cobra"
)

var remoteListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all remotes",
	Long:  "List all configured remote repositories",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Create store
		store, err := internal.NewRemoteStore()
		if err != nil {
			return err
		}

		// Load remotes
		if err := store.Load(); err != nil {
			return err
		}

		// Check if empty
		if len(store.Remotes) == 0 {
			fmt.Println("No remotes configured.")
			return nil
		}

		// Display all remotes
		for _, remote := range store.Remotes {
			fmt.Printf("%s\t%s\n", remote.Name, remote.URL)
		}

		return nil
	},
}

func init() {
	// Register this command to RemoteCmd (parent)
	RemoteCmd.AddCommand(remoteListCmd)
}
