package api

import (
	"encoding/json"
	"net/http"

	"slot-machine-server/internal/auth"

	"slot-machine-server/internal/game"

	"slot-machine-server/internal/payment"
)

func GameSpinHandler(w http.ResponseWriter, r *http.Request) {
	result := game.Spin()
	json.NewEncoder(w).Encode(result)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := auth.Login(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	var deposit struct {
		UserID int     `json:"user_id"`
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&deposit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := payment.Deposit(deposit.UserID, deposit.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
