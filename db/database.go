package db

import (
	"database/sql"
	"goreads/utils"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func GetDb() *sql.DB {
	return db
}

func Init() {
	var err error
	db, err = sql.Open("sqlite", "goreads.db")

	utils.PanicIfEff(err, "No database connection")

	db.SetMaxOpenConns(10)
	prepBooksDbTable()
}

func prepBooksDbTable() {
	createSql := `
	CREATE TABLE IF NOT EXISTS books(
		id 						INTEGER PRIMARY KEY AUTOINCREMENT,
		title 				TEXT,
		isbn 					TEXT,
		author 				TEXT,
		release_year 	INTEGER
	)
	`

	_, err := db.Exec(createSql)

	utils.PanicIfEff(err, "Cannot create table books")
}
