package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=secret dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var nome string
	err = db.QueryRow("SELECT nome FROM produtos WHERE id = $1", 1).Scan(&nome)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nome:", nome)
}
