package controller

//This package is used to band url and method that in service package.

import (
	"GODemo/16_hello_tantan/service"
	"github.com/gorilla/mux"
)

// Bind some url and function which are related to Users
func Users(r *mux.Router){
	r.HandleFunc("/users",service.GetUsers).Methods("GET")
	r.HandleFunc("/users",service.AddUser).Methods("POST")
}