package handler

import (
	"github.com/hellicopthecat/catchlot/users/service"
)

type UserHandler struct {
	userService service.UserService
}

func InitUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
