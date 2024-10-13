package models

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type SpinResult struct {
	Symbols []string `json:"symbols"`
	Win     float64  `json:"win"`
}

type Transaction struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
