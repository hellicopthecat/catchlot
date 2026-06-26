package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v3"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// file:test.db?cache=shared&mode=memory
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatalf("❌ Database is Not Opened :: %d", err)
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) {

	})

	log.Fatalln(app.Listen(":3000"))

}
