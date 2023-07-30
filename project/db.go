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
		rows.Scan(&t.ID, &t.TITLE, &t.DESC, &t.DATE)
		tasks = append(tasks, t)
	}
  return tasks
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
  row.Scan(&t.ID, &t.TITLE, &t.DESC, &t.DATE)
  fmt.Println(t)
  return t

}


func createTask(title string, description string) Task{
	db, _ := sql.Open("sqlite3", "todo.db")

	defer db.Close()
  res, err := db.Exec("insert into tasks(title, description, date)values($1, $2 ,datetime('now'))", title, description)
  if err != nil{
    log.Fatal(err)
  }
 defer db.Close() 
  id, _ := res.LastInsertId()
  row := db.QueryRow("select * from tasks where id=$1", id)
  t := Task{}
  row.Scan(&t.ID, &t.TITLE, &t.DESC, &t.DATE)
  return t


}
