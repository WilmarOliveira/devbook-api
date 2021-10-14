package controller

import (
	"api/src/connection"
	"api/src/entities"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find users"))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user by id"))
}

func InsertUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseErr(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user entities.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.ResponseErr(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.ResponseErr(w, http.StatusBadRequest, err)
		return
	}

	db, err := connection.Connect()
	if err != nil {
		responses.ResponseErr(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.ResponseErr(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseJSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
