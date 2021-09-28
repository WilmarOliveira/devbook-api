package controller

import "net/http"

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find users"))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user by id"))
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
