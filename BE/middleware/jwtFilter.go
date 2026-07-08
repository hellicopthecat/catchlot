package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
)

func JwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_ACCESS_JWT_KEY"))},
		Extractor: extractors.Chain(
			// extractors.FromHeader("Authorization"),
			extractors.FromCookie("ACCESS_TOKEN"),
		),
	})
}
