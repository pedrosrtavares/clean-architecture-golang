package handler

import (
	"clean-architecture-golang/internal/dto/response"
	"clean-architecture-golang/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func MakeUserHandler(r *gin.Engine, service *service.UserService) {
	h := &UserHandler{
		service: service,
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	rg := r.Group("/api/v1")

	rg.GET("/users", h.GetUsers)
	rg.POST("/users")
	rg.PUT("/users/:id")
	rg.DELETE("/users/:id")
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	id := c.Query("id")
	email := c.Query("email")
	if id == "" && email == "" {
		c.AbortWithStatus(400)
		return
	}
	var res response.GetUser
	if id != "" {
		data, err := h.service.GetByID(id)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err := copier.Copy(&res, &data); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, res)
		return
	}
	if email != "" {
		data, err := h.service.GetByEmail(email)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err := copier.Copy(&res, &data); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, res)
		return
	}
}
