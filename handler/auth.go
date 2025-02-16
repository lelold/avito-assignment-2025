package handler

import (
	"avito-assignment-2025/repository"
	"avito-assignment-2025/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func AuthHandler(c *gin.Context) {
	var req struct {
		Name     string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Неверный формат запроса"})
		return
	}

	userRepo := repository.NewUserRepo()
	userService := service.NewUserService(userRepo)

	token, err := userService.Authenticate(req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
