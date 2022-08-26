package main

import (
	"database/sql"
	"log"
	"net/http"
	"user-service/db/models"
)

const webPort = ":8002"

type Config struct {
	DB     *sql.DB
	Models models.Models
}

func main() {
	log.Println("Starting User Service")
	router := CreateRouter()

	log.Fatal(http.ListenAndServe(webPort, router))
}

/*
// Migrating automatically upon starting up
func migrate() {
	log.Println("Migrating to database..")

	exec.Command("cd", "..")
	exec.Command("cd", "..")
	exec.Command("cd", "db")
	exec.Command("cd", "migrations")
	cmd := exec.Command("tern", "migrate")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
*/