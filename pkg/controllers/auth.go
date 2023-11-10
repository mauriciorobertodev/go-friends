package controllers

import (
	"encoding/json"
	"go-friends/pkg/database"
	"go-friends/pkg/models"
	"go-friends/pkg/repositories"
	"go-friends/pkg/responses"
	"go-friends/pkg/security"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	savedUser, err := repository.FindUserByEmail(user.Email)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = security.Verify(savedUser.Password, user.Password)

	if err != nil {
		responses.Json(w, http.StatusBadRequest, err)
		return
	}

	w.Write([]byte("Voce está logado..."))
}
