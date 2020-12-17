package global

// This package is used to store some global variables

import "database/sql"

// To store the connection of db
var Db *sql.DB
