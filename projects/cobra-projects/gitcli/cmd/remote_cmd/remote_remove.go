package remote_cmd

import (
	"fmt"

	"github.com/SHIVANSHGARG07/gitcli/internal"
	"github.com/spf13/cobra"
)

var remoteRemoveCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a remote",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		name := args[0]

		// 1. Create store
		store, err := internal.NewRemoteStore()
		if err != nil {
			return err
		}

		// 2. Load remotes
		if err := store.Load(); err != nil {
			return err
		}

		// 3. Remove remote (in memory)
		if err := store.Remove(name); err != nil {
			return err
		}

		// 4. Save to file
		if err := store.Save(); err != nil {
			return err
		}

		fmt.Printf("Remote '%s' removed\n", name)
		return nil
	},
}

func init() {
	RemoteCmd.AddCommand(remoteRemoveCmd)
}
