package handler

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/hellicopthecat/catchlot/users/request"
)

func (h UserHandler) HCreateUser(c fiber.Ctx) error {
	ctx := context.Background()
	email := c.Query("email")
	social := c.Query("social")
	nickname := c.Query("nickname")

	id, err := uuid.NewV7()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var newUser request.CreateUserRequest
	newUser.Id = id.String()
	newUser.Email = email
	newUser.Social = social
	newUser.Nickname = nickname

	if err := h.userService.SCreateUser(ctx, newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}
