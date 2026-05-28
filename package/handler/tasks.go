package handler

import (
	"main/package/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TasksHandler struct {
	auth service.AuthService
}

func (t *TasksHandler) Action(c *gin.Context) {
	c.JSON(http.StatusOK, "mock")
}
