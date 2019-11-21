package customers

import (
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCustomers(res http.ResponseWriter, req *http.Request) {
	var customer model.Customers
	err := json.NewDecoder(req.Body).Decode(&customer)
	util.OnErr(err)
	db, err := util.ConnectDB()
	util.OnErr(err)
	query := "insert into customers values ('','" + customer.Email + "','" + customer.Name + "','" + customer.Adress + "','" + customer.Phone + "')"
	db.Exec(query)
	util.ReturnRes(res, customer.Name)
}

func GetAllCustomers(res http.ResponseWriter, req *http.Request) {
	var customer model.Customers
	var customer_list []model.Customers
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from customers")
	util.OnErr(err)
	for rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Email, &customer.Name, &customer.Adress, &customer.Phone)
		util.OnErr(err)
		customer_list = append(customer_list, customer)
	}
	util.ReturnRes(res, customer_list)
}

func GetCustomers(res http.ResponseWriter, req *http.Request) {
	var customer model.Customers
	raw_param := mux.Vars(req)
	param := raw_param["id"]
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from customers where id = '" + param + "'")
	util.OnErr(err)
	for rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Email, &customer.Name, &customer.Adress, &customer.Phone)
		util.OnErr(err)
	}
	util.ReturnRes(res, customer)
}

func UpdateCustomers(res http.ResponseWriter, req *http.Request) {
	var customer model.Customers
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	err := json.NewDecoder(req.Body).Decode(&customer)
	db, err := util.ConnectDB()
	util.OnErr(err)
	query := "update customers set email='" + customer.Email + "',name='" + customer.Name + "',adress='" + customer.Adress + "',phone='" + customer.Phone + "' where id='" + id + "'"
	_, err = db.Query(query)
	util.OnErr(err)
	util.ReturnRes(res, customer.Name)
}

func DeleteCustomers(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectDB()
	util.OnErr(err)
	_, err = db.Query("delete from customers where id='" + id + "'")
	util.OnErr(err)
	util.ReturnRes(res, nil)
}
