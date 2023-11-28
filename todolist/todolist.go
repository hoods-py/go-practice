package todolist

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var tasks []string

func init() {
	loadTasks()
	fmt.Println("todolist package initialized with tasks:", tasks)
}

func loadTasks() {
    content, err := ioutil.ReadFile("tasks.txt")
    if err == nil {
        tasks = strings.Split(string(content), "\n")
    }
}


func SaveTasks() {
    content := strings.Join(tasks, "\n")
    err := ioutil.WriteFile("tasks.txt", []byte(content), 0644)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
    }
}


func AddTask(task string) {
	tasks = append(tasks, task)
	fmt.Println("Task added:", task)
}

func ViewTasks() {
	fmt.Println("Tasks before viewing:", tasks)
	fmt.Println("Tasks:")
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func CompleteTask(taskNumber string) {
	index, err := strconv.Atoi(taskNumber)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[index-1] = "[x] " + tasks[index-1][4:]
}

func DeleteTask(taskNumber string) {
	index, err := strconv.Atoi(taskNumber)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
}
