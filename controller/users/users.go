package users

import (
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	var users model.Users
	raw_param := mux.Vars(req)
	param := raw_param["id"]
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from users where id = '" + param + "'")
	util.OnErr(err)
	for rows.Next() {
		err := rows.Scan(&users.Id, &users.Name, &users.Username, &users.Password, &users.Role)
		util.OnErr(err)
	}
	util.ReturnRes(res, users)
}

func GetAllUser(res http.ResponseWriter, req *http.Request) {
	var users model.Users
	var user_list []model.Users
	db, err := util.ConnectDB()
	util.OnErr(err)
	rows, err := db.Query("select * from users")
	util.OnErr(err)
	for rows.Next() {
		err := rows.Scan(&users.Id, &users.Name, &users.Username, &users.Password, &users.Role)
		util.OnErr(err)
		user_list = append(user_list, users)
	}
	util.ReturnRes(res, user_list)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	var user model.Users
	err := json.NewDecoder(req.Body).Decode(&user)
	util.OnErr(err)
	db, err := util.ConnectDB()
	util.OnErr(err)
	query := "insert into users values ('','" + user.Name + "','" + user.Username + "','" + user.Password + "','" + user.Role + "')"
	db.Exec(query)
	util.ReturnRes(res, user.Name)
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	var user model.Users
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	err := json.NewDecoder(req.Body).Decode(&user)
	db, err := util.ConnectDB()
	util.OnErr(err)
	query := "update users set name='" + user.Name + "',username='" + user.Username + "',password='" + user.Password + "',role='" + user.Role + "' where id='" + id + "'"
	log.Println(query)
	_, err = db.Query(query)
	util.OnErr(err)
	util.ReturnRes(res, user.Name)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectDB()
	util.OnErr(err)
	_, err = db.Query("delete from users where id='" + id + "'")
	util.OnErr(err)
	util.ReturnRes(res, nil)
}
