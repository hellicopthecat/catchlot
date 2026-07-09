package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/hellicopthecat/catchlot/gakSoo/repo"
	"github.com/hellicopthecat/catchlot/middleware"
	"github.com/hellicopthecat/catchlot/sqls"
	"github.com/hellicopthecat/catchlot/tickets"
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
	ctx := context.Background()
	gaksoo, err := repo.InitGakSoo(db).RFindGakSooAllID(ctx)

	if err != nil {
		log.Fatalln("❌ GakSoo 초기화 실패")
	}

	// SERVER
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://accounts.google.com"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	api := app.Group("/api")
	jm := middleware.JwtMiddleware()

	users.InitModule(db).UserGroupApi(api, jm)
	tickets.InitTickeModule(db).TicketGroupApi(api, jm, gaksoo)

	log.Fatalln(app.Listen(":4000"))
}
