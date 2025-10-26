package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BarakaCaleb/orderflow/internal/auth"
	"github.com/BarakaCaleb/orderflow/internal/db"
	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/auth/register", auth.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/auth/reset-password", auth.ResetPasswordHandler).Methods("POST")

	port := os.Getenv("PORT")
	log.Printf("Server running on port %s...", port)
	http.ListenAndServe(":"+port, router)
}
