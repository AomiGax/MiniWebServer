package main

// The entrance of the program

import (
	"GODemo/16_hello_tantan/controller"
	"GODemo/16_hello_tantan/global"
	"GODemo/16_hello_tantan/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// main function
// 1. Initial PostgreSQL and create a new db connection for program
// 2. Initial web server, set the port and Handlers
func main() {
	// Initial PostgreSQL and create a new db connection for program
	// global.DB is a global variable. It is used to operation the database
	global.Db = utils.ConnectDB()

	// Use mux to create a new Handlers manager
	r := mux.NewRouter()

	// bind the routes and methods
	controller.Users(r)
	controller.Relationships(r)

	// Specify port and Router
	log.Fatal(http.ListenAndServe(":8088", r))
}
