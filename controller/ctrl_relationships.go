package controller

import (
	"GODemo/16_hello_tantan/service"
	"github.com/gorilla/mux"
)

// Bind some url and function which are related to Relationships
func Relationships(r *mux.Router){
	r.HandleFunc("/users/{uid}/relationships",service.GetRelationships).Methods("GET")
	r.HandleFunc("/users/{uid}/relationships/{other_id}",service.ChangeRelationship).Methods("PUT")
}
