package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func dbConnect() *sql.DB {
	datname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")

	if datname == "" {
		os.Setenv("PGDATABASE", "tasks")
	}

	if sslmode == "" {
		os.Setenv("PGSSLMODE", "disable")
	}

	conn, err := sql.Open("postgres", "")
	checkErr(err)

	return conn
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
