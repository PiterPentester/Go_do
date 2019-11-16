package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "math/rand"
  "strconv"
)

type Task struct {
  ID string `json:"id"`
  Description string `json:"task"`
  Status bool `json:"done"`	
}

var tasks []Task

func getTasks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var task Task
  _ = json.NewDecoder(r.Body).Decode(&task)
  task.ID = strconv.Itoa(rand.Intn(1000000))
  tasks = append(tasks, task)
  json.NewEncoder(w).Encode(&task)
}

func getTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range tasks {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Task{})
}

func updateTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range tasks {
    if item.ID == params["id"] {
      tasks = append(tasks[:index], tasks[index+1:]...)
      var task Task
      _ = json.NewDecoder(r.Body).Decode(&task)
      task.ID = params["id"]
      tasks = append(tasks, task)
      json.NewEncoder(w).Encode(&task)
      
      return
    }
  }
  json.NewEncoder(w).Encode(tasks)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range tasks {
    if item.ID == params["id"] {
      tasks = append(tasks[:index], tasks[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(tasks)
}

func main() {
  fmt.Println("It works!!!")
  router := mux.NewRouter()	
  
  tasks = append(tasks, Task{ID: "1", Description: "My first Task", Status: false})
  
  router.HandleFunc("/tasks", getTasks).Methods("GET")
  router.HandleFunc("/tasks", createTask).Methods("POST")
  router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
  router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
  router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
  
  http.ListenAndServe(":9876", router)
}
