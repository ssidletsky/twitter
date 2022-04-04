package entities

import "time"

// Tweet is a tweet entity
type Tweet struct {
	ID              uint32    `json:"id"`
	Username        string    `json:"username"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Text            string    `json:"text"`
	PublicationDate time.Time `json:"publication_date"`
}
