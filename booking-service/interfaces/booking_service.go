package interfaces

import "booking-service/model"

type IBookingService interface {
	ReserveBooking(venueId, seatId uint) (*model.Booking, error)
	ConfirmBooking(booking *model.Booking) *model.Booking
}
