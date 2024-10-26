package models

type Venue struct {
	ID          uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Location    string `json:"location" gorm:"not null"`
	Seats       []Seat `json:"seats" gorm:"foreignkey:VenueID"`
}
