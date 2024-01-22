package routes

import (
	"github.com/gorilla/mux"
	"github.com/faiznurullah/task-5-pbi-btpns-Faiznurullah/controllers/photocontroller"
	"github.com/faiznurullah/task-5-pbi-btpns-Faiznurullah/middleware"
)

func PhotoRoutes(router *mux.Router) {
	photo := router.PathPrefix("/photos").Subrouter()


	photo.Use(middleware.Auth)
	photo.HandleFunc("", photocontroller.ListPhoto).Methods("GET")
	photo.HandleFunc("", photocontroller.CreatePhoto).Methods("POST")
	photo.HandleFunc("/{photoId}", photocontroller.ShowDetailPhoto).Methods("GET")
	photo.HandleFunc("/{photoId}", photocontroller.UpdatePhoto).Methods("PUT")
	photo.HandleFunc("/{photoId}", photocontroller.DeletePhoto).Methods("DELETE")
}
