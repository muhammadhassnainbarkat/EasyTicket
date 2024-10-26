package service

import (
	"booking-service/model"
	"booking-service/service/interfaces"
)

type BookingService struct {
}

func NewBookingService() interfaces.IBookingService {
	return &BookingService{}
}

func (bookingService BookingService) ReserveBooking(venueId, seatId uint) *model.Booking {
	//TODO implement me
	panic("implement me")
}

func (bookingService BookingService) ConfirmBooking(booking *model.Booking) *model.Booking {
	//TODO implement me
	panic("implement me")
}
