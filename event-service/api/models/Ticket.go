package models

import "time"

type TicketStatus string

const (
	AVAILABLE TicketStatus = "AVAILABLE"
	BOOKED    TicketStatus = "BOOKED"
)

type Ticket struct {
	ID           uint         `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name         string       `json:"name" gorm:"not null"`
	EventID      string       `json:"eventId"`
	TicketStatus TicketStatus `json:"ticketStatus" gorm:"default:'AVAILABLE';type:TicketStatus"`
	ReservedAt   time.Time    `json:"reservedAt"`
	Event        Event        `json:"event"`
	SeatID       uint         `json:"seatId"`
	Seat         Seat         `json:"seat"`
	Price        float64      `json:"price"`
}
