package main

import (
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
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

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_ACCESS_JWT_KEY"))},
		Extractor: extractors.Chain(
			extractors.FromHeader("Authorization"),
		),
	}))
	api := app.Group("/api")

	users.InitModule(db).UserGroupApi(api)

	log.Fatalln(app.Listen(":4000"))

}
