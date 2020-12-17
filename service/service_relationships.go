package service

import (
	"GODemo/16_hello_tantan/dao"
	"GODemo/16_hello_tantan/global"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// The functions about relationships are in this file

// To get someone's all relationship in the "relationships" table
func GetRelationships(w http.ResponseWriter,r *http.Request){
	// Set JSON format
	w.Header().Set("Content-Type","application/json")

	// Return status code
	w.WriteHeader(http.StatusOK)

	// Get the parameters in the url
	params := mux.Vars(r)
	id,_:=strconv.Atoi(params["uid"])

	// Get all relationship for someone
	relationships := dao.GetAllRelationship(global.Db,id)

	// Return JSON to web client
	json.NewEncoder(w).Encode(relationships)
}

// To change the relationship state if two people like each other
func ChangeRelationship(w http.ResponseWriter,r *http.Request){
	// Set JSON format
	w.Header().Set("Content-Type","application/json")

	// Return status code
	w.WriteHeader(http.StatusOK)

	// Get the parameters in the url
	params := mux.Vars(r)
	uid,_:=strconv.Atoi(params["uid"])
	otherID,_:=strconv.Atoi(params["other_id"])

	// Get two people state
	stateA := dao.GetRelationshipState(global.Db,uid,otherID)
	stateB := dao.GetRelationshipState(global.Db,otherID,uid)

	// Compare the state, and update the database
	if stateA=="like" && stateB == "like"{
		dao.ChangeRelationshipState(global.Db,uid,otherID,"match")
		dao.ChangeRelationshipState(global.Db,otherID,uid,"match")
	}
}
