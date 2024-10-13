package payment

import (
	"slot-machine-server/internal/db"
)

func Deposit(userID int, amount float64) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, $3)",
		userID, amount, "deposit")
	if err != nil {
		return err
	}

	return tx.Commit()
}
