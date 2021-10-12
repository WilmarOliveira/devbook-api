package controller

import (
	"api/src/connection"
	"api/src/entities"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatal(err)
	}

	var user entities.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := connection.Connect()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Inserted ID: %d", userID)))

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
