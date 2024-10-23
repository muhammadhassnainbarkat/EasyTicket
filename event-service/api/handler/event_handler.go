package handler

import (
	"event-service/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EventHandler struct {
	EventService models.EventService
}

type EventHandlerConfig struct {
	R            *gin.Engine
	EventService models.EventService
}

func NewEventHandler(cfg *EventHandlerConfig) {
	h := &EventHandler{
		EventService: cfg.EventService,
	}
	g := cfg.R.Group("/api/event")
	g.GET("/ping", h.Ping)
	g.GET("/:id", h.getEventById)
}

func (h *EventHandler) Ping(c *gin.Context) {
	c.JSON(200, h.EventService.Ping())
}
func (h *EventHandler) getEventById(c *gin.Context) {
	id := c.Param("id")
	atoi, _ := strconv.Atoi(id)
	event, err := h.EventService.GetEvent(uint(atoi))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(200, event)
	}
}
