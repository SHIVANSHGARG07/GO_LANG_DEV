package remote_cmd

import (
	"github.com/spf13/cobra"
)

// RemoteCmd is the parent command for all remote subcommands
var RemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage remote repositories",
	Long:  `Add, list, or remove remote repositories.`,
}

// No init() here - children will register themselves to RemoteCmd
// RemoteCmd will be registered to rootCmd from cmd/root.go
