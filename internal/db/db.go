package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func Initialize() error {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"))

	var err error
	DB, err = sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}

	return DB.Ping()
}
