package carts

import (
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//CreateCarts CreateCarts
func CreateCarts(res http.ResponseWriter, req *http.Request) {
	var cart model.Carts
	err := json.NewDecoder(req.Body).Decode(&cart)
	util.OnErr(err)
	db, err := util.ConnectDB()
	util.OnErr(err)
	invoice := string(strconv.Itoa(rand.Intn(1000)) + time.Now().Format("/01/Jan/2006/150405"))
	query := "insert into carts values ('','" + invoice + "','" + strconv.Itoa(cart.Customer_id) + "','" + strconv.Itoa(cart.Total) + "','','')"
	db.Exec(query)
	util.ReturnRes(res, cart.Id)
}

//GetAllCarts Category:Carts
func GetAllCarts(res http.ResponseWriter, req *http.Request) {
	var cart model.Carts
	var carts []model.Carts
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from carts")
	for rows.Next() {
		err := rows.Scan(&cart.Id, &cart.Invoice, &cart.Customer_id, &cart.Total, &cart.Created_At, &cart.Updated_At)
		util.OnErr(err)
		carts = append(carts, cart)
	}
	util.ReturnRes(res, carts)
}

//GetCarts GetCarts
func GetCarts(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	var cart model.Carts
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from carts where id='" + id + "'")
	util.OnErr(err)
	for rows.Next() {
		err := rows.Scan(&cart.Id, &cart.Invoice, &cart.Customer_id, &cart.Total, &cart.Created_At, &cart.Updated_At)
		util.OnErr(err)
	}
	util.ReturnRes(res, cart)
}

//UpdateCarts UpdateCarts
func UpdateCarts(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	var cart model.Carts
	err := json.NewDecoder(req.Body).Decode(&cart)
	util.OnErr(err)
	db, err := util.ConnectDB()
	util.OnErr(err)
	query := "update carts set total='" + strconv.Itoa(cart.Total) + "',updated_at='' where id='" + id + "'"
	log.Println(query)
	_, err = db.Query(query)
	util.OnErr(err)
	util.ReturnRes(res, nil)
}

//DeleteCarts DeleteCarts
func DeleteCarts(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectDB()
	util.OnErr(err)
	_, err = db.Query("delete from carts where id='" + id + "'")
	util.OnErr(err)
	util.ReturnRes(res, nil)
}
