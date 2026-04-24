package handlers

import "net/http"

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User has been created"))
}
