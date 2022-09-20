package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Sample struct {
	ID   int
	Name string
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/sample?charset=utf8&parseTime=true")
	if err != nil {
		log.Printf("failed to open a db err = %s", err.Error())
		return
	}

	if err := db.PingContext(context.Background()); err != nil {
		log.Printf("failed to ping err = %s", err.Error())
		return
	}

	id, err := db.ExecContext(context.Background(), "INSERT INTO sample (name) VALUES (?);", "sample name")
	if err != nil {
		log.Printf("failed to exec query err = %s", err.Error())
		return
	}

	lastid, err := id.LastInsertId()
	if err != nil {
		log.Printf("failed to get a last insert id err = %s", err.Error())
		return
	}

	var s Sample

	if err := db.QueryRowContext(context.Background(), "SELECT id, name FROM sample WHERE id = ?", lastid).Scan(&s.ID, &s.Name); err != nil {
		log.Printf("failed to scan err = %s", err.Error())
		return
	}

	log.Printf("Sample is %#v", s)
}
