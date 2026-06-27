package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/sqls"
)

func main() {
	// init DB
	sqls.InitDB()

	// SERVER
	app := fiber.New()

	api := app.Group("/api")

	api.Get("/", func(c fiber.Ctx) {

	})

	log.Fatalln(app.Listen(":3000"))

}
