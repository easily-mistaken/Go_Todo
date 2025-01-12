package main

import (
	"fmt"
	"net/http"
)
var shortTerm = "Learn Golang"
var longTerm = "Learn Kubernetes"
var reward = "Buy a new laptop"
var taskItems = []string{shortTerm, longTerm, reward}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)

}

func helloUser(w http.ResponseWriter, r *http.Request) {
	var greeting = "Hello User!. Welcome to my Todolist App!"
	fmt.Fprintf(w, "%s", greeting)
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintf(w, "%s\n", task)
	}
}

func printTasks(tasks []string) {
	fmt.Println("List of my Todos :")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func addTask(tasks[]string, newTask string) []string {
	tasks = append(tasks, newTask)
	return tasks
}

func removeTask(tasks[]string, taskIndex int) []string {
	tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
	return tasks
}

func updateTask(tasks[]string, taskIndex int, newTask string) []string {
	tasks[taskIndex] = newTask
	return tasks
}

func completeTask(tasks[]string, taskIndex int) []string {
	tasks = removeTask(tasks, taskIndex)
	return tasks
}

func showTask(tasks[]string, taskIndex int) string {
	return tasks[taskIndex]
}