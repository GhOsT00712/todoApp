package main

import (
	"fmt"
	"net/http"

	"todo"
)

func main() {
	startServer()
}

// function to handle server startup
func startServer() {
	r := http.NewServeMux()
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you have requested path:  %s\n", r.URL.Path)
	})

	http.HandleFunc("POST /Task", addTask)
	http.HandleFunc("DELETE /Task/{taskId}", removeTask)
	http.HandleFunc("GET /Task/{taskId}", getTask)
	http.HandleFunc("GET /Tasks", getAllTask)

	http.ListenAndServe(":80", r)
}

func addTask(w http.ResponseWriter, r *http.Request) {

}

func removeTask(w http.ResponseWriter, r *http.Request) {

}

func getTask(w http.ResponseWriter, r *http.Request) {

}

func getAllTask(w http.ResponseWriter, r *http.Request) {
	tasksIds, err := todo.GetInstance().GetAllTasks()
	if err != nil {
		fmt.Fprintf(w, "Error Occured\n")
	}
	for _, v := range tasksIds {
		fmt.Fprintf(w, "TaskId : %s\n", v)
	}
}
