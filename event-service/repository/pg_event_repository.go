package repository

import (
	"event-service/api/models"
	"fmt"
	"gorm.io/gorm"
)

type PgEventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &PgEventRepository{DB: db}
}

func (repository PgEventRepository) FindByID(u uint) (*models.Event, error) {
	var event models.Event
	repository.DB.First(&event, u)
	return &event, nil
}

func (repository PgEventRepository) CreateEvent(event *models.Event) (*models.Event, error) {
	tx := repository.DB.Create(&event)
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(event)
	return event, nil
}
