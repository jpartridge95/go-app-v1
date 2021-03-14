package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jpartridge95/go-app-v1/environment"
)

// Will be developed alongside handlers to provide functionality to API

func ConnectionOpen() *sql.DB {
	USERNAME := environment.GetVar("username")
	PASSWORD := environment.GetVar("password")
	PORT := environment.GetVar("port")

	db, err := sql.Open("mysql", USERNAME+":"+PASSWORD+"@tcp("+PORT+")/test")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ConnectionClose(db *sql.DB) {
	db.Close()
}
