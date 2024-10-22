package models

type Event struct {
	ID          uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	VenueID     uint      `json:"venueId"`
	Venue       Venue     `json:"venue"`
	PerformerId uint      `json:"performerId"`
	Performer   Performer `json:"performer"`
}
