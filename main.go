package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  //"database/sql"
  //_"github.com/go-sql-driver/mysql"
  //"io/ioutil"
)

type Task struct {
  ID string `json:"id"`
  Description string `json:"task"`
  Status bool `json:"done"`	
}

var tasks []Task

//var db *sql.DB



func main() {
  fmt.Println("It works!!!")
  router := mux.NewRouter()
  
  configureApp()
  
  /*db, err := sql.Open("mysql", "{db_user}:{db_pass}@({db_host}:{db_port})/{db_db}?parseTime=true")
  if err != nil {
    log.Fatal(err)
  }
  if err := db.Ping(); err != nil {
    log.Fatal(err)
  }*/
  
  tasks = append(tasks, Task{ID: "1", Description: "My first Task", Status: false})
  
  router.HandleFunc("/api/v1/tasks", getTasks).Methods("GET")
  router.HandleFunc("/api/v1/tasks", createTask).Methods("POST")
  router.HandleFunc("/api/v1/tasks/{id}", getTask).Methods("GET")
  router.HandleFunc("/api/v1/tasks/{id}", updateTask).Methods("PUT")
  router.HandleFunc("/api/v1/tasks/{id}", deleteTask).Methods("DELETE")
  
  
  http.ListenAndServe(":"+ config["app_port"], router)
}
