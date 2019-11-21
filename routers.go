package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "math/rand"
  "strconv"
  //"database/sql"
  //_"github.com/go-sql-driver/mysql"
  //"io/ioutil"
)

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
