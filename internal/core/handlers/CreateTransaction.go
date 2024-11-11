package handlers

import (
	"fmt"
	"net/http"
	"wallet/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (s *WalletHandler) CreateTransaction(c *gin.Context) {
	var transaction domain.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
	}

	fmt.Printf("Recibido: %v", transaction)
}
