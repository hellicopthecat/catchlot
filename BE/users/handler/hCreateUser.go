package handler

import (
	"context"
	"net/url"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/hellicopthecat/catchlot/users/repo"
	"github.com/hellicopthecat/catchlot/users/request"
	"github.com/hellicopthecat/catchlot/users/service"
)

var statusId string

func HCreateUser(c fiber.Ctx, userRepo repo.UserRepo) error {
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

	if err := service.SCreateUser(ctx, userRepo, newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func HLoginWithGoogle(c fiber.Ctx) error {
	id := uuid.New()
	statusId = id.String()

	q := url.Values{}

	q.Set("client_id", "")
	q.Set("redirect_uri", "") // redirect url
	q.Set("response_type", "code")
	q.Set("scope", "")
	q.Set("state", statusId)

	loginUrl := "https://accounts.google.com/o/oauth2/v2/auth?" + q.Encode()
	return c.Redirect().To(loginUrl)
}
