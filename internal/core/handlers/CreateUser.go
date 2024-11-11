package handlers

import (
	"fmt"
	"net/http"
	"wallet/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (s *WalletHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
	}

	err := s.service.EncryptPassword(&user)
	if err != nil {
		fmt.Println("Error al encryptar contraseña: ", err)
	}

	err = s.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario."})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Usuario creado exitosamente!"})
	}
}
