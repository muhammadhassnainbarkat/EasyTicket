package service

import (
	"booking-service/interfaces"
	"booking-service/model"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type BookingService struct {
	SeatService interfaces.ISeatService
	RedisClient *redis.Client
}

func NewBookingService(SeatService interfaces.ISeatService, RedisClient *redis.Client) interfaces.IBookingService {
	return &BookingService{SeatService, RedisClient}
}

func (bookingService BookingService) ReserveBooking(venueId, seatId uint) (*model.Booking, error) {

	/*
		1. validate the seat
		2. reserve the seat
		3. generate sse to let other know that seat is reserved.
	*/

	seat, err := bookingService.SeatService.FindBySeatAndVenueID(seatId, venueId)
	if err != nil {
		fmt.Println("Seats is not available")
		return nil, errors.New("seats is not available")
	}
	fmt.Println(fmt.Sprintf("Seat: %v ", &seat))

	// 2. reserve the seat.
	/*
		save the seat record in redis for x amount of time
	*/
	locked := bookingService.RedisClient.SetNX(context.Background(), fmt.Sprintf("%v", seatId), seatId, 10*time.Minute)
	if locked.Err() != nil {
		fmt.Println("Error locking booking")
		return nil, errors.New("error locking booking")
	}

	ttl := bookingService.RedisClient.TTL(context.Background(), fmt.Sprintf("%v", seatId))
	result, _ := ttl.Result()
	fmt.Print(fmt.Sprintf("Session Locked for seatId:%d, remaining time: %d\n", seatId, int(result.Seconds())))

	// TODO: generate sse
	return &model.Booking{SeatId: seatId, ExpireInSeconds: int(result.Seconds())}, nil

}

func (bookingService BookingService) ConfirmBooking(booking *model.Booking) *model.Booking {
	//TODO implement me
	panic("implement me")
}
