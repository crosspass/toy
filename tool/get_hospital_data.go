package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "user=root host=/var/run/postgresql dbname=dating password=root123")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select * from hospitals")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
}
