package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "taskcli",
	Short: "A simple task manager CLI",
	Long: `TaskCLI is a command-line task manager that helps you 
manage your daily tasks efficiently.`,
}

// The function already does all the work.
// The only thing it returns is:
// nil
// or an error

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
