package main

import (
	"fmt"
	"go-practice/todolist"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [command] [args]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		todolist.AddTask(strings.Join(args, " "))
		fmt.Println("Task added")
	case "view":
		todolist.ViewTasks()
	case "complete":
		if len(args) < 1 {
			fmt.Println("Usage: go run main.go complete [task number]")
			os.Exit(1)
		}
		taskNumber := args[0]
		todolist.CompleteTask(taskNumber)
		fmt.Println("Task marked as completed")
	case "delete":
		if len(args) < 1 {
			fmt.Println("Usage: go run main.go delete [task number]")
			os.Exit(1)
		}
		taskNumber := args[0]
		todolist.DeleteTask(taskNumber)
		fmt.Println("Task deleted")
	default:
		fmt.Println("Unknown command. Available commands: add, view, complete, delete")
	}
	todolist.SaveTasks()
}
