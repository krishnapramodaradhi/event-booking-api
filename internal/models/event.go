package models

import "time"

type Event struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	TicketCount int       `json:"ticketCount"`
	ImageUrl    string    `json:"ImageUrl"`
	Author      string    `json:"author"`
}
