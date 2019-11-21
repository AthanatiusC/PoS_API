package orders

import (
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllOrders GetAllOrders
func GetAllOrders(res http.ResponseWriter, req *http.Request) {
	var order model.Orders
	var orders []model.Orders
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from orders")
	for rows.Next() {
		rows.Scan(&order.Id, &order.Cart_id, &order.Qty, &order.Price, &order.Created_At)
		orders = append(orders, order)
	}
	util.ReturnRes(res, orders)
}

//CreateOrders CreateOrders
func CreateOrders(res http.ResponseWriter, req *http.Request) {
	var order model.Orders
	err := json.NewDecoder(req.Body).Decode(&order)
	util.OnErr(err)
	db, err := util.ConnectDB()
	util.OnErr(err)
	query := "insert into orders values (''," + strconv.Itoa(order.Cart_id) + "," + strconv.Itoa(order.Qty) + "," + strconv.Itoa(order.Price) + ",'')"
	db.Exec(query)
	util.ReturnRes(res, order.Id)
}

func GetOrders(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	var order model.Orders
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from orders where id='" + id + "'")
	for rows.Next() {
		err := rows.Scan(&order.Id, &order.Cart_id, &order.Qty, &order.Price, &order.Created_At)
		util.OnErr(err)
	}
	util.ReturnRes(res, order)
}
