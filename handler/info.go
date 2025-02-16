package handler

import (
	"avito-assignment-2025/repository"
	"avito-assignment-2025/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoHandler(c *gin.Context) {
	user, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}

	infoService := service.NewInfoService(
		repository.NewUserRepo(),
		repository.NewTransactionRepo(),
		repository.NewItemRepo(),
		repository.NewBuyRepo(),
	)

	info, err := infoService.GetInfo(user.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, info)
}
