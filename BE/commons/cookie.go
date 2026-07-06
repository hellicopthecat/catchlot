package commons

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/constants"
)

func RegistCookies(c fiber.Ctx, email string, name string, social string) (*string, *string) {
	at, err := GenerateACCESSJWT(email, social, name)
	if err != nil {
		TokenError(c, err)
		return nil, nil
	}
	rt, err := GenerateREFRESHJWT()
	if err != nil {
		TokenError(c, err)
		return nil, nil
	}
	RegistCookie(c, constants.ACCESS, at)
	RegistCookie(c, constants.REFRESH, rt)
	return nil, &rt
}

func RegistCookie(c fiber.Ctx, name string, token string) {
	cookie := new(fiber.Cookie)

	var max int

	if name == constants.ACCESS {
		max = 60 * 60 * 24 * 7
	} else {
		max = 60 * 60 * 24
	}

	cookie.Name = name
	cookie.HTTPOnly = true
	cookie.Path = "/"
	cookie.Secure = true
	cookie.SameSite = "Lax"
	cookie.MaxAge = max
	cookie.Value = token

	c.Cookie(cookie)
}
