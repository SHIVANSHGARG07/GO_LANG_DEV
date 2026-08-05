package internal

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskStore struct {
	Tasks    []Task `json:"tasks"`
	filepath string
}

// this creates. new task store
// returns taskstore
func NewTaskStore(filepath string) *TaskStore {
	return &TaskStore{
		Tasks:    []Task{}, // Start with empty list
		filepath: filepath, // Remember where to save
	}
}

// load read from file
func (ts *TaskStore) Load() error {
	data, err := os.ReadFile(ts.filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	// takes json and convert it into taskstore type
	return json.Unmarshal(data, ts)
}

// save write to file
func (ts *TaskStore) Save() error {
	data, err := json.MarshalIndent(ts, "", " ")

	if err != nil {
		return err
	}

	// returns error
	return os.WriteFile(ts.filepath, data, 0644)
}

// add. new task
func (ts *TaskStore) Add(title, description string) {
	id := len(ts.Tasks) + 1 // New ID = current count + 1
	task := Task{           // Create new task
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	ts.Tasks = append(ts.Tasks, task) // Add to list
}

// complete mark task done
func (ts *TaskStore) Complete(id int) bool {
	for i := range ts.Tasks { // Loop through tasks
		if ts.Tasks[i].ID == id { // Found the right one?
			ts.Tasks[i].Completed = true // Mark it done
			return true                  // Success!
		}
	}
	return false // Didn't find it
}

// delete task
func (ts *TaskStore) Delete(id int) bool {
	for i := range ts.Tasks {
		if ts.Tasks[i].ID == id {
			// Remove by slicing: [before] + [after]
			ts.Tasks = append(ts.Tasks[:i], ts.Tasks[i+1:]...)
			return true
		}
	}
	return false
}
