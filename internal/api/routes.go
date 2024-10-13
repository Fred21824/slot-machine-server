package api

import (
	"slot-machine-server/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/auth/login", LoginHandler).Methods("POST")

	// Protected routes
	s := r.PathPrefix("/api").Subrouter()
	s.Use(middleware.AuthMiddleware)
	s.HandleFunc("/game/spin", GameSpinHandler).Methods("POST")
	s.HandleFunc("/payment/deposit", DepositHandler).Methods("POST")

	return r
}
