package handler

import (
	"github.com/gin-gonic/gin"
	todo_rest_api "github.com/mTeeeur/todo-rest-api"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCtx)

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	var input todo_rest_api.TodoList

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) getAllLists(c *gin.Context) {}

func (h *Handler) getListById(c *gin.Context) {}

func (h *Handler) updateList(c *gin.Context) {}

func (h *Handler) deleteList(c *gin.Context) {}
