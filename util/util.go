package util

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func ConnectDB() (database *sql.DB, err error) {
	db, err := sql.Open("mysql", "root@tcp(localhost)/lsp")
	return db, err
}

func ReturnRes(res http.ResponseWriter, value interface{}) {
	res.Header().Set("content-type", "application/json")
	json.NewEncoder(res).Encode(value)
}

func OnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
