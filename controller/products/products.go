package products

import (
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateProduct Category:Products
func CreateProduct(res http.ResponseWriter, req *http.Request) {
	var product model.Products
	err := json.NewDecoder(req.Body).Decode(&product)
	util.OnErr(err)
	db, err := util.ConnectDB()
	util.OnErr(err)
	query := "insert into product values ('','" + product.Name + "'," + strconv.Itoa(product.Price) + "," + strconv.Itoa(product.Stock) + ")"
	db.Exec(query)
	util.ReturnRes(res, product.Name)
}

//GetProduct GetProduct
func GetProduct(res http.ResponseWriter, req *http.Request) {
	var product model.Products
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from product where id='" + id + "'")
	util.OnErr(err)
	for rows.Next() {
		err := rows.Scan(rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock))
		util.OnErr(err)
	}
	util.ReturnRes(res, product)
}

//GetAllProduct GetAllProduct
func GetAllProduct(res http.ResponseWriter, req *http.Request) {
	var product model.Products
	var products []model.Products
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from product")
	for rows.Next() {
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
		products = append(products, product)
	}
	util.ReturnRes(res, products)
}

//UpdateProduct UpdateProduct
func UpdateProduct(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	var product model.Products
	err := json.NewDecoder(req.Body).Decode(&product)
	util.OnErr(err)
	db, err := util.ConnectDB()
	util.OnErr(err)
	query := "update product set name='" + product.Name + "',price=" + strconv.Itoa(product.Price) + ",stock=" + strconv.Itoa(product.Stock) + " where id='" + id + "'"
	log.Println(query)
	_, err = db.Query(query)
	util.OnErr(err)
	util.ReturnRes(res, nil)
}

//DeleteProduct DeleteProduct
func DeleteProduct(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectDB()
	util.OnErr(err)
	_, err = db.Query("delete from product where id ='" + id + "'")
	util.OnErr(err)
	util.ReturnRes(res, nil)
}
