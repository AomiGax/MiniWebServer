package dao

import (
	"GODemo/16_hello_tantan/entity"
	"database/sql"
	"fmt"
	"log"
)

// Get someone's all relationship
func GetAllRelationship(db *sql.DB, uid int) []entity.Relationship {
	// Declare the variables for relationships
	var id, otherID int
	var state string

	// Declare a array to put all relationships
	var relationshipArr []entity.Relationship

	// Get all relationships from database
	rows, err := db.Query(" select id,other_id,state from public.relationships where uid=$1", uid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Assign the variables and put the relationship into array
	for rows.Next() {
		err := rows.Scan(&id, &otherID, &state)

		if err != nil {
			log.Fatal(err)
		}
		relationshipArr = append(relationshipArr, entity.Relationship{ID: id, UID: uid, OtherID: otherID, State: state, Type: "relationship"})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return relationshipArr
}

// Update someone's one relationship into some other state
func ChangeRelationshipState(db *sql.DB, uid int, otherID int, state string) {
	stmt, err := db.Prepare("UPDATE relationships set state=$1 WHERE uid=$2 and other_id=$3")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(state, uid, otherID)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("update relationship_tbl success")
	}
}

// Get state of someone's relationship
func GetRelationshipState(db *sql.DB, uid int, otherID int) string {
	var state string
	row := db.QueryRow(" select state from public.relationships where uid=$1 and other_id=$2", uid, otherID)
	err := row.Scan(&state)
	if err != nil {
		log.Fatal(err)
	}
	return state
}
