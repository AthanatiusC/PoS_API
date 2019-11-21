package main

import (
	controller_carts "LSP/controller/cart"
	controller_customers "LSP/controller/customers"
	controller_orders "LSP/controller/orders"
	controller_products "LSP/controller/products"
	controller_users "LSP/controller/users"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Users struct {
	Id       int
	Name     string
	Username string
	Password string
	Role     string
}

type Products struct {
	Id    int
	Name  string
	Price int
	Stock int
}

type Customers struct {
	Id     int
	Email  string
	Name   string
	Adress string
	Phone  string
}

type Carts struct {
	Id          int
	Invoice     string
	Customer_id int
	Total       int
}

type Orders struct {
	Id         int
	Cart_id    int
	Qty        int
	Price      int
	Created_At string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")

	router.HandleFunc("/users", controller_users.GetAllUser).Methods("GET")
	router.HandleFunc("/users", controller_users.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controller_users.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controller_users.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controller_users.DeleteUser).Methods("DELETE")

	router.HandleFunc("/products", controller_products.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controller_products.GetAllProduct).Methods("GET")
	router.HandleFunc("/products/{id}", controller_products.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", controller_products.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controller_products.DeleteProduct).Methods("DELETE")

	router.HandleFunc("/orders", controller_orders.GetAllOrders).Methods("GET")
	router.HandleFunc("/orders", controller_orders.CreateOrders).Methods("POST")
	router.HandleFunc("/orders/{id}", controller_orders.GetOrders).Methods("GET")
	// router.HandleFunc("/orders/{id}", UpdateOrders).Methods("PUT")
	// router.HandleFunc("/orders/{id}", DeleteOrders).Methods("DELETE")

	router.HandleFunc("/carts", controller_carts.CreateCarts).Methods("POST")
	router.HandleFunc("/carts", controller_carts.GetAllCarts).Methods("GET")
	router.HandleFunc("/carts/{id}", controller_carts.GetCarts).Methods("GET")
	router.HandleFunc("/carts/{id}", controller_carts.UpdateCarts).Methods("PUT")
	router.HandleFunc("/carts/{id}", controller_carts.DeleteCarts).Methods("DELETE")

	router.HandleFunc("/customers", controller_customers.CreateCustomers).Methods("POST")
	router.HandleFunc("/customers", controller_customers.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controller_customers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controller_customers.UpdateCustomers).Methods("PUT")
	router.HandleFunc("/customers/{id}", controller_customers.DeleteCustomers).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:3030", router))
}

//Home Display Category:Default
func Home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Go lang REST API</h1><br><span>Made with love by Lexi Anugrah</span>"))
}
