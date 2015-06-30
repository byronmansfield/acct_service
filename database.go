package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgress"
	DB_NAME     = "accounts"
)

var currentId int

var todos Todos

func init() {
	RepoCreatedTodo(Todo{Name: "Write presentation"})
	RepoCreatedTodo(Todo{Name: "Host meetup"})
}

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s pasword=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postres", dbinfo)
	PanicIf(err)
	defer db.Close()
}

func PanicIf(err, error) {
	if err != nil {
		panic(err)
	}
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func RepoCreatedTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
