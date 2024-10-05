package server

import (
	"fmt"
	"log"
	"net/http"
	"wallet/internals/core/router"
)

func InitServer(port uint) {
	route := router.InitRouter()

	fmt.Printf("Server is running at port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), route))
}
