package cmd

import (
	"fmt"
	"github.com/SHIVANSHGARG07/taskcli/internal"
	"github.com/spf13/cobra"
	"strconv"
)

var deleteCmd = &cobra.Command{

	Use:   "delete",
	Short: "Delete task",
	Long:  "Delete task by ID",

	RunE: func(cmd *cobra.Command, args []string) error {

		// string to int
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("failed in conversion from string to int", args[0])
		}

		// load tasks
		store := internal.NewTaskStore("tasks.json")
		if err := store.Load(); err != nil {
			return err
		}

		// modify task.json
		if !store.Delete(id) {
			return fmt.Errorf("task with ID %d not found", id)
		}

		// save
		if err := store.Save(); err != nil {
			return err
		}

		// print message
		fmt.Printf("Task %d deleted successfully", id)
		return nil
	},
}

func init() {

	rootCmd.AddCommand(deleteCmd)

}
