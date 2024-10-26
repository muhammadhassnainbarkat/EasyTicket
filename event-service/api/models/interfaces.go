package models

type EventService interface {
	Ping() map[string]any
	GetEvent(uint) (*Event, error)
	CreateEvent(event *Event) (*Event, error)
	GetAllEvents(page int, size int) []Event
}

type EventRepository interface {
	FindByID(uint) (*Event, error)
	CreateEvent(event *Event) (*Event, error)
	FindAllEvents(page, size int) []Event
}
