package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(cfg *Config) {
	h := &Handler{}

	g := cfg.R.Group("/api/events")
	g.GET("/ping", h.Ping)
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
