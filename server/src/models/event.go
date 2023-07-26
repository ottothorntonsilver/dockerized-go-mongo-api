package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string
	VenueID     string `bson:"venue_id,omitempty"`
	Date        string `bson:"date,omitempty"`
	TicketCount int    `bson:"ticket_count,omitempty"`
}
