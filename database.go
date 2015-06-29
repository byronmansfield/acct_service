package main

import (
	"fmt"
	"github.com/demandbase/acct_service/models"
)

var currentId int

var todos Todos

func init() {
	RepoCreatedTodo(Todo{Name: "Write presentation"})
	RepoCreatedTodo(Todo{Name: "Host meetup"})
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
