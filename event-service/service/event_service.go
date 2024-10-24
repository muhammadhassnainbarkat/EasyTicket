package service

import "event-service/api/models"

type EventService struct {
	EventRepository models.EventRepository
}

type ESConfig struct {
	EventRepository models.EventRepository
}

func NewEventService(config *ESConfig) models.EventService {
	return &EventService{
		EventRepository: config.EventRepository,
	}
}

func (es *EventService) GetEvent(ID uint) (*models.Event, error) {
	return &models.Event{
		ID:          ID,
		Name:        "Test",
		Description: "Description",
	}, nil
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
