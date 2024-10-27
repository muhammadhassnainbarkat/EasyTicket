package service

import (
	"booking-service/interfaces"
	"booking-service/model"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type BookingService struct {
	SeatService interfaces.ISeatService
	RedisClient *redis.Client
}

func NewBookingService(SeatService interfaces.ISeatService, RedisClient *redis.Client) interfaces.IBookingService {
	return &BookingService{SeatService, RedisClient}
}

func (bookingService BookingService) ReserveBooking(venueId, seatId uint) *model.Booking {

	/*
		1. validate the seat
		2. reserve the seat
		3. generate sse to let other know that seat is reserved.
	*/

	_, err := bookingService.SeatService.FindBySeatAndVenueID(seatId, venueId)
	if err != nil {
		fmt.Println("Seats is not available")
		return nil
	}

	// 2. reserve the seat.
	/*
		save the seat record in redis for x amount of time
	*/

	// TODO: 3. generate the sse event to let other know that seat has been reserved.

	return &model.Booking{}

}

func (bookingService BookingService) ConfirmBooking(booking *model.Booking) *model.Booking {
	//TODO implement me
	panic("implement me")
}
