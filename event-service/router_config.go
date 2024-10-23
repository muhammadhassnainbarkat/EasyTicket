package main

import (
	"event-service/api/handler"
	"event-service/repository"
	"event-service/service"
	"github.com/gin-gonic/gin"
)

func setupRouter(d *DataSources) (*gin.Engine, error) {

	router := gin.New()

	eventRepository := repository.NewEventRepository(d.DB)
	eventService := service.NewEventService(&service.ESConfig{EventRepository: eventRepository})

	h := handler.EventHandlerConfig{R: router, EventService: eventService}
	handler.NewEventHandler(&h)

	return router, nil
}
