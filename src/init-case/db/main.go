package main

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

/*
* init
(1) Environment var
(2) The DB connection is stored to 'db', which is a global var.

init doesn't return any errors.
The only way to notice error is panic.
And panic occurs, application crashes.

Secondly its hard to test.
If init activates first and testcases activated,
it may go against the original intent.

Thirdly we had to save db connection pool to global var.
 * if any func in same package, it can change the global var.
 * We can't seperate func which dependent to global var,
   so its hard to write unit tests.
*/

func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME") // (1)
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	db = d // (2)
}

func createClient(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
