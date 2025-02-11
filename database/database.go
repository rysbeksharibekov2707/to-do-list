package database

import (
	"database/sql"
	"log"
	"todoactual/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg *config.Config) error {
	var err error
	DB, err = sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		return err
	}
	log.Println("Connected to the database")
	return DB.Ping()
}
