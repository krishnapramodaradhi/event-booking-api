package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/krishnapramodaradhi/event-booking-api/internal/config"
	"github.com/krishnapramodaradhi/event-booking-api/internal/models"
	"github.com/krishnapramodaradhi/event-booking-api/internal/util"
)

var events = []models.Event{
	{
		Id:          "123-abc",
		Title:       "Test Event",
		Description: "Test Description",
		Price:       123.123,
		StartTime:   time.Now(),
		EndTime:     time.Now().AddDate(0, 0, 4),
		Author:      "me",
	},
	{
		Id:          "1234-abc",
		Title:       "Test Event",
		Description: "Test Description",
		Price:       1234.1234,
		StartTime:   time.Now(),
		EndTime:     time.Now().AddDate(0, 0, 2),
		Author:      "me",
	},
}

func FetchEvents(w http.ResponseWriter, r *http.Request, db *config.Database) (int, error) {
	return http.StatusOK, util.WriteJSON(w, http.StatusOK, events)
}

func FetchEventById(w http.ResponseWriter, r *http.Request, db *config.Database) (int, error) {
	params := mux.Vars(r)
	for _, event := range events {
		if event.Id == params["id"] {
			return http.StatusOK, util.WriteJSON(w, http.StatusOK, event)
		}
	}
	return http.StatusNotFound, fmt.Errorf("id: %v is not found", params["id"])
}

// func handleCreateEvents(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

// func handleUpdateEvents(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

// func handleDeleteEvents(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
