package main

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-http-app/todo"
)

var todoList = todo.ToDoList{FilePath: "todos.json"}

func main() {
	todoList.Load()
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", handleTasksById)
	http.ListenAndServe(":8080", nil)
}

func handleTasks(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		for _, task := range todoList.Tasks {
			fmt.Fprintf(writer, "%d - %s\n", task.Id, task.Title)
		}
	case "POST":
		title := request.URL.Query().Get("title")
		todoList.Add(title)
		fmt.Fprintf(writer, "Task added: %s", title)
	}
}

func handleTasksById(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "DELETE" {
		stringId := request.URL.Path[len("/tasks/"):]
		id, _ := strconv.Atoi(stringId)
		todoList.Remove(id)
		fmt.Fprintf(writer, "Task with Id: %d has been removed", id)
	}
}
