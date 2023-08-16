package router

import (
	"github.com/4rivappa/S20200010022/second-question/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/numbers", controller.GetNumbers).Methods("GET")

	return router
}
