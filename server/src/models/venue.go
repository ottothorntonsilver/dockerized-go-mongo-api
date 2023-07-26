package models

type Venue struct {
	Name     string
	Address  interface{} `bson:"address,omitempty"`
	Capacity int         `bson:"capacity,omitempty"`
}
