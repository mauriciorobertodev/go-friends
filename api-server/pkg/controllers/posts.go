package controllers

import (
	"encoding/json"
	"errors"
	"go-friends/pkg/authentication"
	"go-friends/pkg/database"
	"go-friends/pkg/models"
	"go-friends/pkg/repositories"
	"go-friends/pkg/responses"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StorePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(body, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := post.Prepare(true); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorId = userId

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db)
	post.Id, err = repository.Store(post)

	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusCreated, post)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	post, err := repository.FindPostById(postId)

	if err != nil {
		responses.Error(w, http.StatusNotFound, err)
		return
	}

	responses.Json(w, http.StatusOK, post)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	users, err := repository.Search(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, users)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != id {
		responses.Error(w, http.StatusForbidden, errors.New("you cannot update post of another user"))
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(body, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(false); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	err = repository.UpdatePost(id, post)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != id {
		responses.Error(w, http.StatusForbidden, errors.New("you cannot delete post of another user"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	err = repository.DeletePost(id)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusNoContent, nil)
}

func GetPostsOfUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	users, err := repository.GetPostsOfUser(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, users)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId == id {
		responses.Error(w, http.StatusForbidden, errors.New("you cannot like your own post"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	err = repository.LikePost(id)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusNoContent, nil)
}
