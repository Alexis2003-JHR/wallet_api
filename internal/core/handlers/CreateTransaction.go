package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *WalletHandler) CreateTransaction(c *gin.Context) {
	fmt.Println("Works")
}
