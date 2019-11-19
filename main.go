package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

type Task struct {
  ID string `json:"id"`
  Description string `json:"task"`
  Status bool `json:"done"`	
}

var tasks []Task

func main() {
  fmt.Println("It works!!!")
  router := mux.NewRouter()
  
  db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/mysql?parseTime=true")
  if err != nil {
    log.Fatal(err)
  }
  if err := db.Ping(); err != nil {
    log.Fatal(err)
  }
  
  tasks = append(tasks, Task{ID: "1", Description: "My first Task", Status: false})
  
  router.HandleFunc("/api/v1/tasks", getTasks).Methods("GET")
  router.HandleFunc("/api/v1/tasks", createTask).Methods("POST")
  router.HandleFunc("/api/v1/tasks/{id}", getTask).Methods("GET")
  router.HandleFunc("/api/v1/tasks/{id}", updateTask).Methods("PUT")
  router.HandleFunc("/api/v1/tasks/{id}", deleteTask).Methods("DELETE")
  
  http.ListenAndServe(":9876", router)
}
