package models

type EventService interface {
	Ping() map[string]any
	GetEvent(uint) (*Event, error)
}

type EventRepository interface {
	FindByID(uint) (*Event, error)
}
