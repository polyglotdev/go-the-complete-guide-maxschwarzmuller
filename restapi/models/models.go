package models

import "time"

// Event represents the data model for an event entity.
type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

// Save saves the event to the database.
func (e *Event) Save() error {
	// Save the event to the database...
	events = append(events, *e)
	return nil
}
