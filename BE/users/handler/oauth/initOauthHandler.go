package oauth

import "github.com/hellicopthecat/catchlot/users/service"

type OauthHandler struct {
	userService service.UserService
}

func InitOauthHandler(service service.UserService) *OauthHandler {
	return &OauthHandler{
		userService: service,
	}
}
