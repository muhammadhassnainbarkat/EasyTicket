package model

type Booking struct {
	SeatId          uint `json:"seatId"`
	ExpireInSeconds int  `json:"expireInSeconds"`
}
