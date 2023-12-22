package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Database struct {
	db *sql.DB
}

func NewDB() (*Database, error) {
	dbUrl := fmt.Sprintf("%v?authToken=%v", os.Getenv("DB_URL"), os.Getenv("AUTH_TOKEN"))
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("connected to db")
	return &Database{
		db: db,
	}, nil
}
