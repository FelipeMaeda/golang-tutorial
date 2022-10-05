package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	Name string
	Done bool
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/tasks", TaskHandler)
	http.ListenAndServe(":8888", nil)
}

func Hello(web http.ResponseWriter, r *http.Request) {
	web.Write([]byte("Olá mundo"))
}

func TaskHandler(web http.ResponseWriter, r *http.Request) {
	task := Task{
		Name: "Aprender Go",
		Done: true,
	}
	jayson, _ := json.Marshal(task)
	web.Write(jayson)
	// t := template.Must(template.ParseFiles("task.html"))
	// t.Execute(web, task)
}
