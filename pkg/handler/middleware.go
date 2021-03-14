package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) UserIdentify(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}
