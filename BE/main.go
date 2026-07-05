package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/sqls"
	"github.com/hellicopthecat/catchlot/users"
)

func main() {
	// init DB
	db := sqls.InitDB()
	defer db.Close()

	// SERVER
	app := fiber.New()

	api := app.Group("/api")

	users.InitModule(db).UserGroupApi(api)

	log.Fatalln(app.Listen(":3000"))

}
