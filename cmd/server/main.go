package main

import (
	"database/sql"
	"github.com/Sclea3/Pet-simple-HTTP-server/internal/db"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}(database)
}
