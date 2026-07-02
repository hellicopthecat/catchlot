package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/sqls"
)

func main() {
	// init DB
	db := sqls.InitDB()

	// SERVER
	app := fiber.New()

	api := app.Group("/api")

	log.Fatalln(app.Listen(":3000"))

}
