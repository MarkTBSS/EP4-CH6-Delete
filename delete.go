package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	database, err := sql.Open("postgres", "user=postgres password=Pass1234 host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer database.Close()
	stagement, err := database.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		log.Fatal("can't prepare delete statement", err)
	}
	if _, err := stagement.Exec(1); err != nil {
		log.Fatal("can't execute delete statment", err)
	}
	fmt.Println("delete success")
}
