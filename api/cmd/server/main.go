package main

import (
	"log"
	"net/http"
	"orderflow/internal/auth"
	"orderflow/internal/db"

	"github.com/gorilla/mux"
)

func main() {
	db.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/auth/register", auth.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", auth.LoginHandler).Methods("POST")

	log.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", router)
}
