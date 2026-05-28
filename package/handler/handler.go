package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	handler http.Handler
}

func (h *Handler) ServeRoutes() {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/signup")
		auth.POST("/signin")
		auth.POST("/logout")
	}
	tasks := router.Group("/tasks")
	{
		tasks.GET("/")
		tasks.GET("/:id")
		tasks.POST("/")
		tasks.PATCH("/:id")
		tasks.DELETE("/:id")
	}
}
