package router

import (
	"wallet/internals/core/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HandleGenerateData)

	return r
}
