package handler

import (
	"github.com/gin-gonic/gin"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
)

func (h *Handler) login(c *gin.Context) {
	var input todo_rest_api.User

	if err := c.BindJSON(&input); err != nil {
	}
}

func (h *Handler) register(c *gin.Context) {
}
