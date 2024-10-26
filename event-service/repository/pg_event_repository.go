package repository

import (
	"event-service/api/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgEventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &PgEventRepository{DB: db}
}

func (repository PgEventRepository) FindByID(u uint) (*models.Event, error) {
	var event models.Event
	if err := repository.DB.Preload("Venue.Seats").Preload(clause.Associations).First(&event, u).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
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

func (repository PgEventRepository) FindAllEvents(page, size int) []models.Event {
	var events []models.Event
	repository.DB.Preload("Venue.Seats").Preload(clause.Associations).Scopes(Paginate(page, size)).Find(&events)
	return events
}

func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
