package cmd

import (
	"fmt"
	"github.com/SHIVANSHGARG07/taskcli/internal"
	"github.com/spf13/cobra"
)

var showCompleted bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `Display all tasks in your task list.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		// load tasks file
		store := internal.NewTaskStore("tasks.json")
		if err := store.Load(); err != nil {
			return err
		}

		if len(store.Tasks) == 0 {
			fmt.Println("No tasks found. Add one with 'taskcli add -t \"Your task\"'")
			return nil
		}

		fmt.Println("\nYour Tasks:")
		fmt.Println("===========")
		for _, task := range store.Tasks {
			if !showCompleted && task.Completed {
				continue
			}

			status := "[ ]"
			if task.Completed {
				status = "[✓]"
			}

			fmt.Printf("%s ID: %d - %s\n", status, task.ID, task.Title)
			if task.Description != "" {
				fmt.Printf("    %s\n", task.Description)
			}
			fmt.Printf("    Created: %s\n\n", task.CreatedAt.Format("2006-01-02 15:04"))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showCompleted, "all", "a", true, "Show completed tasks")
}
