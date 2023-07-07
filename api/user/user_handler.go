package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func Newhandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.Service.CreateUser(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
