package controllers

import (
	"encoding/json"
	"fmt"
	"go-friends/pkg/database"
	"go-friends/pkg/models"
	"go-friends/pkg/repositories"
	"io"
	"log"
	"net/http"
)

func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários..."))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário..."))
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	id, err := repository.Store(user)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", id)))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário..."))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário..."))
}
