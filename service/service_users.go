package service

// This package id used to complete some function or logic

import (
	"GODemo/16_hello_tantan/dao"
	"GODemo/16_hello_tantan/entity"
	"GODemo/16_hello_tantan/global"
	"encoding/json"
	"net/http"
)

// The functions about users are in this file

// To get all users in the "users" table
func GetUsers(w http.ResponseWriter,r *http.Request){
	// Set JSON format
	w.Header().Set("Content-Type","application/json")

	// Return status code
	w.WriteHeader(http.StatusOK)

	// Get all users
	users := dao.GetAllUser(global.Db)

	// Return JSON to web client
	json.NewEncoder(w).Encode(users)
}

// To add a new user into the "users" table
func AddUser(w http.ResponseWriter,r *http.Request){
	// Set JSON format
	w.Header().Set("Content-Type","application/json")

	// Return status code
	w.WriteHeader(http.StatusOK)

	// Get information from request.body, and create a new user
	var user entity.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Get the max id, and assign the new userid
	user.ID = dao.GetUserID(global.Db)+1
	user.Type = "user"

	// Add user into database
	dao.CreateUser(global.Db,user)

	// Return JSON to web client
	json.NewEncoder(w).Encode(user)
}