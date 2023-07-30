package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID    int    `json:"id"`
	TITLE string `json:"title"`
	DESC  string `json:"description"`
	DATE  string `json:"date"`
  STATUS bool  `json:"status"`
}

func getTasks() []Task{
	db, err := sql.Open("sqlite3", "todo.db")
  if err != nil{
    log.Fatal(err)
  }
	defer db.Close()
	rows, err := db.Query("select * from tasks")
  if err != nil{
    log.Fatal(err)
  }
	defer db.Close()

	tasks := []Task{}
	for rows.Next() {
		t := Task{}
		rows.Scan(&t.ID, &t.TITLE, &t.DESC, &t.DATE, &t.STATUS)
		tasks = append(tasks, t)
	}
  return tasks
}


func deleteTasks(){
  db, err := sql.Open("sqlite3", "todo.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()

  db.Exec("delete from tasks")
  defer db.Close()
}

func deleteDetailTask(id int){
  db, err := sql.Open("sqlite3", "todo.db")
  if err != nil{
      log.Fatal(err)
    }
  defer db.Close()
  db.Exec("delete from tasks where id=$1", id)
  defer db.Close()

}
func putDetailTask(id int, status bool){
  db, err := sql.Open("sqlite3", "todo.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()
  db.Exec("update tasks set status=$1 where id=$2", status, id )
  defer db.Close()
}

func getDetailTask(id int) Task{
  db, err := sql.Open("sqlite3", "todo.db")
  if err != nil{
    log.Fatal(err)
  }
  defer db.Close()
  row := db.QueryRow("select * from tasks where id=$1", id)
  defer db.Close()
  fmt.Println(id)
  t := Task{}
  row.Scan(&t.ID, &t.TITLE, &t.DESC, &t.DATE, &t.STATUS)
  fmt.Println(t)
  return t

}


func createTask(title string, description string, status bool) Task{
	db, _ := sql.Open("sqlite3", "todo.db")

	defer db.Close()
  res, err := db.Exec("insert into tasks(title, description, date, status)values($1, $2 ,datetime('now'), $3)", title, description, status)
  if err != nil{
    log.Fatal(err)
  }
 defer db.Close() 
  id, _ := res.LastInsertId()
  row := db.QueryRow("select * from tasks where id=$1", id)
  t := Task{}
  err = row.Scan(&t.ID, &t.TITLE, &t.DESC, &t.DATE, &t.STATUS)
  if err != nil{
    log.Fatal(err)
  }
  return t


}
