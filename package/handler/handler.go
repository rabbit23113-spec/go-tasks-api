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
	authHandler := new(AuthHandler)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", authHandler.SignUp)
		auth.POST("/signin", authHandler.SignIn)
	}
	tasks := router.Group("/tasks")
	{
		tasks.GET("/")
		tasks.GET("/:id")
		tasks.POST("/")
		tasks.PATCH("/:id")
		tasks.DELETE("/:id")
	}
	return router
}
