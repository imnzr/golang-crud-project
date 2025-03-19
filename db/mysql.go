package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/imnzr/golang-crud-project/config"
)

var DB *sql.DB

func ConnectMySQL(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBname,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error opening MySQL connection: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to MySQL database: %v", err)
	}
	log.Println("Successfully connected to MySQL database!")
	DB = db
}
