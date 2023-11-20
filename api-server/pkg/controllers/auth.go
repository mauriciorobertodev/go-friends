package controllers

import (
	"encoding/json"
	"go-friends/pkg/authentication"
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

	token, err := authentication.CreateToken(savedUser.Id)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, token)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(body, &password); err != nil {
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
	oldHashedPassword, err := repository.GetPasswordById(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Verify(oldHashedPassword, password.Actual); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	newHashedPassword, err := security.Hash(password.New)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePasswordById(userId, string(newHashedPassword)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusNoContent, nil)
}
