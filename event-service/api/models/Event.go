package models

type Event struct {
	ID          uint       `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	VenueID     *uint      `json:"venueId" gorm:"default:null"`
	Venue       *Venue     `json:"venue"`
	PerformerId *uint      `json:"performerId" gorm:"default:null"`
	Performer   *Performer `json:"performer"`
}
