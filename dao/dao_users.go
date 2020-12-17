package dao

import (
	"GODemo/16_hello_tantan/entity"
	"database/sql"
	"fmt"
	"log"
)

// This package is used to operate the database

// Get the max userID
func GetUserID(db *sql.DB) int {
	// Declare the maxID
	var max int

	// Get max ID from database
	row := db.QueryRow(" select count(id) from public.users")
	err := row.Scan(&max)
	if err != nil {
		fmt.Println(err)
	}

	return max
}

// Get all users' information
func GetAllUser(db *sql.DB) []entity.User {
	// Declare the variables for user
	var id int
	var name string

	// Declare a array to put all users
	var userArr []entity.User

	// Get all users from database
	rows, err := db.Query(" select id,name from public.\"users\"")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Assign the variables and put the user into array
	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}
		userArr = append(userArr, entity.User{ID: id, Name: name,Type: "user"})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return userArr
}

// Add a new user
func CreateUser(db *sql.DB, user entity.User) {
	stmt, err := db.Prepare("insert into public.\"users\" (id,name) values($1,$2)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.ID, user.Name)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("insert into user_tbl success")
	}
}
