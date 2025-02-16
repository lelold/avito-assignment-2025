package handler

import (
	"avito-assignment-2025/repository"
	"avito-assignment-2025/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendCoinRequest struct {
	ToUser string `json:"toUser" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

func SendCoinHandler(c *gin.Context) {
	var req SendCoinRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Неверный формат запроса"})
		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Сумма перевода должна быть больше 0"})
		return
	}

	fromUserName, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован"})
		return
	}

	userRepo := repository.NewUserRepo()
	txRepo := repository.NewTransactionRepo()
	txService := service.NewTransactionService(userRepo, txRepo)

	err := txService.TransferCoins(fromUserName.(string), req.ToUser, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Транзакция успешна"})
}
