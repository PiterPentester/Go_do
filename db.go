package main

import (
  "fmt"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

func createDB(name string) {

   db, err := sql.Open("mysql", config["db_user"] + ":" + config["db_pass"] + "@tcp(" + config["db_host"] + ":" + config["db_port"] + ")/")
   if err != nil {
       panic(err)
   }

   _, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
   if err != nil {
       panic(err)
   }
   
   _, err = db.Exec("USE " + name)
   if err != nil {
       fmt.Println(err.Error())
   } else {
       fmt.Println("DB selected successfully..")
   }
   
   stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS tasks(id varchar(15) NOT NULL, task varchar(50), done varchar(5), PRIMARY KEY (id));")
   if err != nil {
       fmt.Println(err.Error())
   }
   _, err = stmt.Exec()
   if err != nil {
       fmt.Println(err.Error())
   } else {
       fmt.Println("Table created successfully..")
   }
   defer db.Close()
}


