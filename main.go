package main

import (
	"github.com/imnzr/golang-crud-project/config"
	"github.com/imnzr/golang-crud-project/db"
)

func main() {
	cfg := config.LoadConfig()
	db.ConnectMySQL(cfg)
}
