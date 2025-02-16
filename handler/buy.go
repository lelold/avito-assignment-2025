package handler

import (
	"avito-assignment-2025/repository"
	"avito-assignment-2025/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BuyRequest struct {
	Merch string
}

func BuyHandler(c *gin.Context) {
	itemName := c.Param("item")

	if itemName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Предмет не указан"})
		return
	}

	user, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован"})
		return
	}

	userRepo := repository.NewUserRepo()
	merchRepo := repository.NewItemRepo()
	butRepo := repository.NewBuyRepo()

	purchService := service.NewBuyService(userRepo, merchRepo, butRepo)

	if err := purchService.BuyItem(user.(string), itemName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успещно приобретён"})
}
