package handlers

import (
	"net/http"

	_ "github.com/go-playground/validator/v10"
)

type userRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User has been created"))
}
