package handler

import (
	"github.com/gin-gonic/gin"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
	"net/http"
)

func (h *Handler) login(c *gin.Context) {
	var input todo_rest_api.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) register(c *gin.Context) {
}
