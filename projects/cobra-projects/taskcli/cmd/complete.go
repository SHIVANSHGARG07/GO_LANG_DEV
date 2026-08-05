package cmd

import (
	"fmt"
	"strconv"

	"github.com/SHIVANSHGARG07/taskcli/internal"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task-id]",
	Short: "Mark a task as completed",
	Long:  `Mark a task as completed by its ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// convert string to int
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID: %s", args[0])
		}

		// load tasks
		store := internal.NewTaskStore("tasks.json")
		if err := store.Load(); err != nil {
			return err
		}

		if !store.Complete(id) {
			return fmt.Errorf("task with ID %d not found", id)
		}

		if err := store.Save(); err != nil {
			return err
		}

		fmt.Printf("✓ Task %d marked as completed!\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
