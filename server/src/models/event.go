package models

type Event struct {
	Name        string
	VenueID     string `bson:"venue_id,omitempty"`
	Date        string `bson:"date,omitempty"`
	TicketCount int    `bson:"ticket_count,omitempty"`
}
