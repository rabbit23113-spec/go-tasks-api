package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	handler http.Handler
}

func (h *Handler) ServeRoutes() *gin.Engine {
	router := gin.New()
	AuthHandler := new(AuthHandler)
	TasksHandler := new(TasksHandler)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", AuthHandler.SignUp)
		auth.POST("/signin", AuthHandler.SignIn)
	}
	tasks := router.Group("/tasks")
	{
		tasks.GET("/", TasksHandler.Action)
		tasks.GET("/:id", TasksHandler.Action)
		tasks.POST("/", TasksHandler.Action)
		tasks.PATCH("/:id", TasksHandler.Action)
		tasks.DELETE("/:id", TasksHandler.Action)
	}
	return router
}
