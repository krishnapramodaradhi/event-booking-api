package main

import (
	"log"

	"github.com/krishnapramodaradhi/event-booking-api/internal/config"
	"github.com/krishnapramodaradhi/event-booking-api/internal/server"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func init() {
	config.LoadEnv()
}

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal("error occured while connecting to db", err)
	}
	s := server.New(":8080", db)
	s.Run()
}
