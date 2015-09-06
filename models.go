package main

import (
	"net/http"
	"time"
)

// apiError define structure of API error
type apiError struct {
	Tag     string `json:"-"`
	Error   error  `json:"-"`
	Message string `json:"error"`
	Code    int    `json:"code"`
}

// ApiHandler global API mux
type ApiHandler struct {
	Handler func(w http.ResponseWriter, r *http.Request) *apiError
}

type Route struct {
	Name    string
	Method  string
	Pattern string
	ApiHandler
}

type Routes []Route

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

var todos Todos
