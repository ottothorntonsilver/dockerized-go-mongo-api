package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string
	VenueID     string `bson:"venue_id,omitempty"`
	Date        string `bson:"date,omitempty"`
	TicketCount int    `bson:"ticket_count,omitempty"`
}

type EventSchema struct {
	Name        string `json:"name"`
	VenueID     string `json:"venue_id"`
	Date        string `json:"date"`
	TicketCount int    `json:"ticket_count"`
}
