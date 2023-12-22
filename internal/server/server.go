package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnapramodaradhi/event-booking-api/internal/config"
	"github.com/krishnapramodaradhi/event-booking-api/internal/controllers"
	"github.com/krishnapramodaradhi/event-booking-api/internal/util"
)

type Server struct {
	listenAddress string
	db            *config.Database
}

func New(listenAddress string, db *config.Database) *Server {
	return &Server{
		listenAddress: listenAddress,
		db:            db,
	}
}

func (s *Server) Run() {
	r := mux.NewRouter()
	// events router
	er := r.PathPrefix("/api/events").Subrouter()
	er.HandleFunc("", util.MakeHttpHandleFunc(controllers.FetchEvents, s.db)).Methods("GET")
	er.HandleFunc("/{id}", util.MakeHttpHandleFunc(controllers.FetchEventById, s.db)).Methods("GET")

	log.Println("Server started on port:", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, r))
}
