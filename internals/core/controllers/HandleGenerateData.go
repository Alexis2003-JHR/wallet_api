package controllers

import (
	"fmt"
	"net/http"
	"wallet/internals/application"
	"wallet/internals/core/domain"
	"wallet/internals/infrastructure/postgres"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := postgres.NewPostgresDB()
	if err != nil {
		panic("Error al conectar a la base de datos")
	}
	fmt.Println("Está conectando")
	userRepo := postgres.NewGormRepository(db)
	userService := application.NewUserService(userRepo)
	fmt.Println("Está inicializando")
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
