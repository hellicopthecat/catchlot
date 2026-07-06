package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/hellicopthecat/catchlot/sqls"
	"github.com/hellicopthecat/catchlot/users"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(".env파일이 존재하지 않습니다.")
	}
	// init DB
	db := sqls.InitDB()
	defer db.Close()

	// SERVER
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://accounts.google.com"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))
	api := app.Group("/api")

	users.InitModule(db).UserGroupApi(api)

	log.Fatalln(app.Listen(":4000"))

}
