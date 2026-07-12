package commons

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func BadSQLFile(err error) {
	log.Fatalf("❌ SQL 디렉토리 및 파일 읽기 실패 %s :: ", err)
}

func TokenError(c fiber.Ctx, err error) error {
	log.Printf("❌ Token 생성 실패 %s", err)
	return c.Status(fiber.StatusUnauthorized).JSON(Results{
		Status: false,
		Msg:    "로그인에 실패하였습니다.",
	})
}

func UnauthorizedError(c fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Results{
		Status: false,
		Msg:    "잘못된 인증요청입니다.",
	})
}

func CheckAtoi(c fiber.Ctx, err error) error {
	log.Printf("❌ CheckAtoi Err :: %s", err)
	return c.Status(fiber.StatusBadRequest).JSON(Results{
		Status: false,
		Msg:    "잘못된 요청입니다.",
	})
}
