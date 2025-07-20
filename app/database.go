package app

import (
	"database/sql"
	"golangpzn-restfulapi/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "golang-pzn:Golang-PZN123@tcp(localhost:3306)/golang_pzn_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
