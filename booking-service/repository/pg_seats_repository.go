package repository

import (
	"booking-service/interfaces"
	"booking-service/model"
	"fmt"
	"gorm.io/gorm"
)

type PgSeatsRepository struct {
	DB *gorm.DB
}

func NewPgSeatsRepository(db *gorm.DB) interfaces.ISeatRepository {
	return &PgSeatsRepository{DB: db}
}

func (repository PgSeatsRepository) FindFirstBySeatAndVenueID(seatId, venueId uint) (*model.Seat, error) {
	var seat model.Seat
	tx := repository.DB.Where("id = ? AND venue_id = ?", seatId, venueId).First(&seat)
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(seat)
	return &seat, nil
}
