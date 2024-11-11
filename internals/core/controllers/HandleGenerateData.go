package controllers

import (
	"fmt"
	"wallet/internals/core/domain"
	"wallet/internals/core/services"

	postgres "wallet/cmd/database"
	"wallet/internals/repository/user"

	"github.com/gin-gonic/gin"
)

func HandleCreateUser(c *gin.Context) {
	db, err := postgres.NewPostgresDB()
	if err != nil {
		panic("Error al conectar a la base de datos")
	}

	userRepo := user.NewGormRepository(db)
	userService := services.NewUserService(userRepo)

	user := domain.User{
		FirstName: "Heyder Alexis",
		LastName:  "Rojas Pineda",
		Address:   "Carrera 134 B 15s",
		Email:     "alexisrojas23@gmail.com",
		Age:       23,
	}

	err = userService.CreateUser(user)
	if err != nil {
		fmt.Println("Error al crear usuario: ", err)
	} else {
		fmt.Println("Usuario creado exitosamente!")
	}
}
