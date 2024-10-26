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
	if err := repository.DB.First(&event, u).Error; err != nil {
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
	var count int64
	repository.DB.Scopes(Paginate(page, size)).Find(&events).Count(&count)
	fmt.Println(count)
	fmt.Println(events)
	return events
}

func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
