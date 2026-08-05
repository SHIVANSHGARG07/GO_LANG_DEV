package cmd

import (
	"fmt"
	"github.com/SHIVANSHGARG07/taskcli/internal"
	"github.com/spf13/cobra"
)

var (
	taskTitle       string
	taskDescription string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to your task list with a title and optional description.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if taskTitle == "" {
			return fmt.Errorf("task title is required")
		}

		store := internal.NewTaskStore("tasks.json")
		if err := store.Load(); err != nil {
			return err
		}

		store.Add(taskTitle, taskDescription)

		if err := store.Save(); err != nil {
			return err
		}

		fmt.Printf("✓ Task added: %s\n", taskTitle)
		return nil
	},
}

// Special Go function that runs automatically when package loads

func init() {
	rootCmd.AddCommand(addCmd)

	// Flags

	// taskcli add --title "Buy milk"
	// taskcli add -t "Buy milk"
	addCmd.Flags().StringVarP(&taskTitle, "title", "t", "", "Task title (required)")

	addCmd.Flags().StringVarP(&taskDescription, "description", "d", "", "Task description")

	// cobra show error if title is not provided
	addCmd.MarkFlagRequired("title")
}
