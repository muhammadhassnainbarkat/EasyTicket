package service

import (
	"context"
	"event-service/api/models"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type EventService struct {
	EventRepository models.EventRepository
	RedisClient     *redis.Client
}

type ESConfig struct {
	EventRepository models.EventRepository
	RedisClient     *redis.Client
}

func NewEventService(config *ESConfig) models.EventService {
	return &EventService{
		EventRepository: config.EventRepository,
		RedisClient:     config.RedisClient,
	}
}

func (es *EventService) GetEvent(ID uint) (*models.Event, error) {

	event, err := es.EventRepository.FindByID(ID)

	if err != nil {
		return nil, err
	}

	seats := event.Venue.Seats
	availableSeats := make([]models.Seat, 0, len(seats))
	if len(seats) > 0 {
		seatsID := make([]string, len(seats))
		for i, seat := range seats {
			seatsID[i] = fmt.Sprintf("%v", seat.ID)
		}
		lookedSeats, err := es.RedisClient.MGet(context.Background(), seatsID...).Result()
		if err == nil {
			for i, lookedSeat := range lookedSeats {
				if lookedSeat == nil {
					availableSeats = append(availableSeats, seats[i])
				}
			}
		}
	}
	event.Venue.Seats = availableSeats

	return event, nil
}

func (es *EventService) Ping() map[string]any {
	mapping := make(map[string]any)
	mapping["message"] = "pong"
	return mapping
}

func (es *EventService) CreateEvent(event *models.Event) (*models.Event, error) {
	createEvent, err := es.EventRepository.CreateEvent(event)
	if err != nil {
		return nil, err
	}
	return createEvent, nil
}

func (es *EventService) GetAllEvents(page int, size int) []models.Event {
	return es.EventRepository.FindAllEvents(page, size)
}
