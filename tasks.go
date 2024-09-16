package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func loadTask(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if fileInfo.Size() == 0 {
		return []Task{}, nil
	}
	err = json.NewDecoder(file).Decode(&tasks)
	return tasks, err
}
func saveTask(filename string, tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)

}

func addTasks(tasks []Task, description string) []Task {
	id := len(tasks) + 1
	task := Task{
		ID:          id,
		Description: description,
		Completed:   false,
	}
	return append(tasks, task)
}

func listTasks(tasks []Task, showCompeleted bool) {
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "complete"
		}
		if showCompeleted == task.Completed {
			fmt.Printf("%d. %s [%s]\n", task.ID, task.Description, status)

		}

	}
}
func completeTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
		}
	}
	return tasks
}

func deleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func editTask(tasks []Task, id int, newDescription string) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
		}
	}
	return tasks
}
