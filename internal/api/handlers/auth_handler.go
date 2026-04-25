package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/basel2053/social-media/database"
	"github.com/go-playground/validator/v10"
)

type userRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

var validate = validator.New()

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	user := userRequest{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	existingUser, err := database.Queries.GetUserByEmail(context.Background(), user.Email)
	fmt.Println(existingUser)

	// w.Write([]byte("User has been created"))
}
