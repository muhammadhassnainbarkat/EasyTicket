package models

type Seat struct {
	ID         uint64 `json:"id"`
	SeatNumber string `json:"seatNumber"`
	VenueID    uint64 `json:"venueId"`
}
