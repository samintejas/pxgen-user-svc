package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"pxgen.io/user/internal/repo"
	"pxgen.io/user/internal/rest/request"
	"pxgen.io/user/internal/utils"
)

type AuthHandler struct {
	repo repo.AuthRepositoryInterface
}

func NewAuthHandler(repo repo.AuthRepositoryInterface) *AuthHandler {
	return &AuthHandler{repo: repo}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload request.LoginRequest

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	hash, err := h.repo.GetHashedPassword(payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	ok := utils.ComparePassword(hash, payload.Password)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	token, err := utils.GenerateJWT(payload.Username)

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   true,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}
