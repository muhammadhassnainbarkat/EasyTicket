package repository

import (
	"event-service/api/models"
	"gorm.io/gorm"
)

type PgEventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &PgEventRepository{DB: db}
}

func (p PgEventRepository) FindByID(u uint) (*models.Event, error) {
	var event models.Event
	p.DB.First(&event, u)
	return &event, nil
}
