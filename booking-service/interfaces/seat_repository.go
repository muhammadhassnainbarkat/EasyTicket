package interfaces

import "booking-service/model"

type ISeatRepository interface {
	FindFirstBySeatAndVenueID(seatId, venueId uint) (*model.Seat, error)
}
