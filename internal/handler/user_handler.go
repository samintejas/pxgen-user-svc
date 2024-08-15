package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pxgen.io/user/internal/models"
	"pxgen.io/user/internal/repo"
	"pxgen.io/user/internal/rest/request"
	"pxgen.io/user/internal/rest/response"
	"pxgen.io/user/internal/utils"
)

type UserHandler struct {
	repo repo.UserRepositoryInterface
}

func NewUserHandler(repo repo.UserRepositoryInterface) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)

	if !ok {
		http.Error(w, "Could not get user claims", http.StatusInternalServerError)
		return
	}

	user, err := h.repo.GetUserByUsername(claims.Username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateUser

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	ex, err := h.repo.ExcistsByUsernameAndEmail(payload.UserName, payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	hashedPass, err := utils.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("bad password"))
	}

	if !ex {
		newuser := models.User{
			UserName:  payload.UserName,
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Password:  hashedPass,
			Status:    payload.Status,
		}
		userId, err := h.repo.CreateUser(&newuser)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create new user"))
		} else {
			utils.WriteJson(w, http.StatusCreated, response.RegistedUser{ID: userId})
		}
	} else {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user/email already registered"))
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var payload request.UpdateUser
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	claims, ok := r.Context().Value("claims").(*utils.Claims)

	if !ok {
		http.Error(w, "Could not get user claims", http.StatusInternalServerError)
		return
	}

	u, err := h.repo.GetUserByUsername(claims.Username)
	if u == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	} else {

		updateuser := new(models.User)
		updateuser.ID = u.ID
		updateuser.FirstName = payload.FirstName
		updateuser.LastName = payload.LastName
		updateuser.UserName = payload.UserName
		updateuser.Email = payload.Email
		updateuser.Status = payload.Status
		updateuser.Password = payload.Password

		latestuser, err := h.repo.UpdateUser(updateuser)

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		utils.WriteJson(w, http.StatusOK, latestuser)
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*utils.Claims)

	if !ok {
		http.Error(w, "Could not get user claims", http.StatusInternalServerError)
		return
	}

	if err := h.repo.DeleteUser(claims.Username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
