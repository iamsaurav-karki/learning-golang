package main

import (
	"fmt"
	"net/http"
)

var taskItems []string // global slice to hold tasks

func main() {
	fmt.Println("Welcome to Todo List App:")
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show_tasks", showTasks)
	http.HandleFunc("/add_task", addTaskHandler) // ✅ moved here
	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	if len(taskItems) == 0 {
		fmt.Fprintln(writer, "No tasks yet.")
		return
	}
	for i, task := range taskItems {
		fmt.Fprintf(writer, "%d: %s\n", i+1, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our TodoApp!"
	fmt.Fprintf(writer, greeting)
}

func printTasks(taskItems []string) {
	fmt.Println("List of todos")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := r.URL.Query().Get("task")
	if task == "" {
		http.Error(w, "Task cannot be empty", http.StatusBadRequest)
		return
	}
	taskItems = append(taskItems, task)
	fmt.Fprintf(w, "Added task: %s\n", task)
}
