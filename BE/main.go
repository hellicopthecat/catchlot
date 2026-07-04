package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/sqls"
	"github.com/hellicopthecat/catchlot/users/handler"
	"github.com/hellicopthecat/catchlot/users/repo"
)

func main() {
	// init DB
	db := sqls.InitDB()
	defer db.Close()
	userRepo := repo.InitUserRepo(db)

	// SERVER
	app := fiber.New()

	api := app.Group("/api")

	api.Post("/users", func(c fiber.Ctx) error {
		return handler.HCreateUser(c, userRepo)
	})

	log.Fatalln(app.Listen(":3000"))

}
