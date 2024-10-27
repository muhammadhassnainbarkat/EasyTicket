package service

import (
	"booking-service/interfaces"
	"booking-service/model"
)

type SeatService struct {
	seatsRepository interfaces.ISeatRepository
}

func NewSeatsService(seatsRepository interfaces.ISeatRepository) interfaces.ISeatService {

	return &SeatService{seatsRepository: seatsRepository}
}

func (seatService SeatService) FindBySeatAndVenueID(seatId, venueId uint) (*model.Seat, error) {
	return seatService.seatsRepository.FindFirstBySeatAndVenueID(seatId, venueId)
}
