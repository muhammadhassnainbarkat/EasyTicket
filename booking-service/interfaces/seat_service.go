package interfaces

import "booking-service/model"

type ISeatService interface {
	FindBySeatAndVenueID(seatId, venueId uint) (*model.Seat, error)
}
