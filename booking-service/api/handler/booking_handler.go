package handler

import (
	"booking-service/service/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookingHandler struct {
	BookingService interfaces.IBookingService
}

type ReserveBookingRequest struct {
	venueId uint
	seatId  uint
}

func NewBookingHandler(engine *gin.Engine, bookingService interfaces.IBookingService) {

	handler := &BookingHandler{
		BookingService: bookingService,
	}
	group := engine.Group("/api/booking")
	{
		group.POST("/reserve", handler.reserveBooking)
		group.POST("/confirm", handler.confirmBooking)
	}
}

func (bookingHandler *BookingHandler) reserveBooking(context *gin.Context) {

	request := ReserveBookingRequest{}
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	booking := bookingHandler.BookingService.ReserveBooking(request.venueId, request.seatId)

	context.JSON(http.StatusOK, booking)
}

func (bookingHandler *BookingHandler) confirmBooking(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}
