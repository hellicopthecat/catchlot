package handler

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
)

func (h UserHandler) HLogout(c fiber.Ctx) error {
	ctx := context.Background()
	at := c.Cookies(constants.ACCESS)

	claims, err := commons.DecodedACCESSJWT(at)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(commons.Results{
			Status: false,
			Msg:    "잘못된 인증요청입니다.",
		})
	}

	h.userService.SLogout(ctx, claims.Email)

	c.Cookie(&fiber.Cookie{
		Name:     constants.ACCESS,
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:     constants.REFRESH,
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(commons.Results{
		Status: true,
		Msg:    "로그아웃 성공",
	})
}
