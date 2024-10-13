package auth

import (
	"errors"

	"slot-machine-server/pkg/models"

	"slot-machine-server/internal/db"
)

func Login(username, password string) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRow("SELECT id, username, balance FROM users WHERE username = $1 AND password = $2",
		username, password).Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}
