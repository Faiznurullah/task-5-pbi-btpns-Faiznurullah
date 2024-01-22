package routes

import (
	"github.com/gorilla/mux"
	"github.com/faiznurullah/task-5-pbi-btpns-Faiznurullah/controllers/usercontroller"
	"github.com/faiznurullah/task-5-pbi-btpns-Faiznurullah/middleware"
)

func UserRoutes(router *mux.Router) {
	user := router.PathPrefix("/users").Subrouter()

	user.Use(middleware.Auth)
	user.HandleFunc("", usercontroller.ListUser).Methods("GET")
	user.HandleFunc("/{userId}", usercontroller.UpdateUser).Methods("PUT")
	user.HandleFunc("/{userId}", usercontroller.DeleteUser).Methods("DELETE")
}
