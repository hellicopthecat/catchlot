package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/users/handler"
	"github.com/hellicopthecat/catchlot/users/handler/oauth"
	"github.com/hellicopthecat/catchlot/users/repo"
	"github.com/hellicopthecat/catchlot/users/service"
)

type UserModule struct {
	userHandler  *handler.UserHandler
	oauthHandler *oauth.OauthHandler
}

func InitModule(db *sql.DB) *UserModule {
	userRepo := repo.InitUserRepo(db)
	userService := service.InitUserService(userRepo)
	oauthHandler := oauth.InitOauthHandler(*userService)
	userHandler := handler.InitUserHandler(*userService)
	return &UserModule{
		userHandler:  userHandler,
		oauthHandler: oauthHandler,
	}
}

func (m *UserModule) UserGroupApi(g fiber.Router) {
	u := g.Group("/user")
	u.Get("/google/login", m.oauthHandler.LoginWithGoogleRequest)
	u.Get("/google/response", m.oauthHandler.LoginWithGoogleResponse)
	u.Get("/logout", m.userHandler.HLogout)
}
