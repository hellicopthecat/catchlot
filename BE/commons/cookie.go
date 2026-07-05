package commons

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/constants"
)

func RegistCookies(c fiber.Ctx, email string, name string, social string) {
	at, err := GenerateACCESSJWT(email, social, name)
	if err != nil {
		TokenError(c, err)
	}
	rt, err := GenerateREFRESHJWT()
	if err != nil {
		TokenError(c, err)
	}
	RegistCookie(c, constants.ACCESS, at)
	RegistCookie(c, constants.REFRESH, rt)

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
