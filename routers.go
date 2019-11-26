package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "math/rand"
  "strconv"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "io/ioutil"
  "fmt"
)


func getTasks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  
  db, err := sql.Open("mysql", config["db_user"] + ":" + config["db_pass"] + "@tcp(" + config["db_host"] + ":" + config["db_port"] + ")/")
   if err != nil {
       panic(err)
   }
  defer db.Close()
  
  _, err = db.Exec("USE " + config["db_db"])
   if err != nil {
       fmt.Println(err.Error())
   }
    
  var tasks []Task
    
  result, err := db.Query("SELECT id, task, done FROM tasks")
  if err != nil {
    panic(err.Error())
  }  
  defer result.Close()
  
  for result.Next() {
    var task Task
    err := result.Scan(&task.ID, &task.Description, &task.Status)
    if err != nil {
      panic(err.Error())
    }
    tasks = append(tasks, task)
  }
   
  json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  
  db, err := sql.Open("mysql", config["db_user"] + ":" + config["db_pass"] + "@tcp(" + config["db_host"] + ":" + config["db_port"] + ")/")
  if err != nil {
      panic(err)
  }
  defer db.Close()
  
  _, err = db.Exec("USE " + config["db_db"])
  if err != nil {
      fmt.Println(err.Error())
  }
  
  stmt, err := db.Prepare("INSERT INTO tasks(id, task, done) VALUES(?, ?, ?)")
  if err != nil {
    panic(err.Error())
  }
  
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  id := strconv.Itoa(rand.Intn(1000000))
  task := keyVal["task"]
  done := keyVal["done"]
  
  _, err = stmt.Exec(id, task, done)
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "New task was created")
}

func getTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  
  db, err := sql.Open("mysql", config["db_user"] + ":" + config["db_pass"] + "@tcp(" + config["db_host"] + ":" + config["db_port"] + ")/")
  if err != nil {
      panic(err)
  }
  defer db.Close()
  
  _, err = db.Exec("USE " + config["db_db"])
  if err != nil {
      fmt.Println(err.Error())
  }
  
  params := mux.Vars(r)
  
  result, err := db.Query("SELECT id, task, done FROM tasks WHERE id = ?", params["id"])
  if err != nil {
      panic(err.Error())
  }
  
  var task Task
  for result.Next() {
    err := result.Scan(&task.ID, &task.Description, &task.Status)
    if err != nil {
      panic(err.Error())
    }
  }
    
  
  json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  
  db, err := sql.Open("mysql", config["db_user"] + ":" + config["db_pass"] + "@tcp(" + config["db_host"] + ":" + config["db_port"] + ")/")
  if err != nil {
      panic(err)
  }
  defer db.Close()
  
  _, err = db.Exec("USE " + config["db_db"])
  if err != nil {
      fmt.Println(err.Error())
  }
  
  params := mux.Vars(r)
  
  stmt, err := db.Prepare("UPDATE tasks SET task = ?, done = ? WHERE id = ?")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  
  if err != nil {
    panic(err.Error())
  }
  
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  
  newTask := keyVal["task"]
  newStatus := keyVal["done"]
  
  _, err = stmt.Exec(newTask, newStatus, params["id"])
  
  if err != nil {
    panic(err.Error())
  }
  fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  
  db, err := sql.Open("mysql", config["db_user"] + ":" + config["db_pass"] + "@tcp(" + config["db_host"] + ":" + config["db_port"] + ")/")
  if err != nil {
      panic(err)
  }
  defer db.Close()
  
  _, err = db.Exec("USE " + config["db_db"])
  if err != nil {
      fmt.Println(err.Error())
  }
  
  params := mux.Vars(r)
  
  stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
  if err != nil {
    panic(err.Error())
  }
  
  _, err = stmt.Exec(params["id"])
  if err != nil {
    panic(err.Error())
  }
  
  fmt.Fprintf(w, "Task with ID = %s was deleted", params["id"])
  
}
