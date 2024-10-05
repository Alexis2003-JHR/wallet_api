package controllers

import (
	"fmt"
	"net/http"
)

func HandleGenerateData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Prueba desde Handle de generar data")
}
