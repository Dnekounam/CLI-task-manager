package main

import (
	"fmt"
	"os"
	"strconv"
)

const taskFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-task-manager <command> [arguments]")
		return
	}

	tasks, err := loadTask(taskFile)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch os.Args[1] {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: cli-task-manager add <task description")
			return
		}
		description := os.Args[2]
		tasks = addTasks(tasks, description)
		err = saveTask(taskFile, tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	case "complete":
		if len(os.Args) < 3 {
			fmt.Printf("Usage: cli-task-manager complete <task ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid tasl ID")
			return
		}
		tasks = completeTask(tasks, id)
		err = saveTask(taskFile, tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: cli-task-manager delete <task ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		tasks = deleteTask(tasks, id)
		err = saveTask(taskFile, tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)

		}
	case "list":
		if len(os.Args) > 2 && os.Args[2] == "completed" {
			listTasks(tasks, true)
		} else {
			listTasks(tasks, false)
		}
	case "edit":
		if len(os.Args) < 4 {
			fmt.Println("Usage: cli-task-manager edit <task ID> <new description>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		newDescription := os.Args[3]
		tasks = editTask(tasks, id, newDescription)
		err = saveTask(taskFile, tasks)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
