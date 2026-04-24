package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/basel2053/social-media/internal/api/handlers"
	v1 "github.com/basel2053/social-media/internal/api/routes/v1"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to our API"))
}

func main() {
	port := "8080"

	fmt.Println("Server is up and running on port 8080")
	http.HandleFunc("GET /", handlerRoot)
	http.HandleFunc("GET /healthz", handlerHealth)
	handler := handlers.Handler{}
	v1.RegisterV1Routes(http.DefaultServeMux, &handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
