package routes

import (
	"net/http"
	controller "netLog/Controllers"
)

func SetupRoutes() {

    http.HandleFunc("/", controller.HelloHandler)
    http.HandleFunc("/new-visitor", controller.AddNewVisitor)
    http.HandleFunc("/log", controller.GetUsers)

}
