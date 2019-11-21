package model

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
	Created_At  string
	Updated_At  string
}

type Orders struct {
	Id         int
	Cart_id    int
	Qty        int
	Price      int
	Created_At string
}
