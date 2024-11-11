package controllers

import (
	"fmt"
	"net/http"
	"wallet/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (s *UserHandler) HandleCreateUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
	}

	err := s.service.CreateUser(user)
	if err != nil {
		fmt.Println("Error al crear usuario: ", err)
	} else {
		fmt.Println("Usuario creado exitosamente!")
	}
}
