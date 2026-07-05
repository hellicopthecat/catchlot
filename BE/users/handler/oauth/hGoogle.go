package oauth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/users/request"
)

type googleTokenResponse struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
}

type googleClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func (o OauthHandler) LoginWithGoogleRequest(c fiber.Ctx) error {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")

	if clientId == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(commons.Results{
			Status: false,
			Msg:    "서버에 문제가 생겼습니다.",
		})
	}

	stateId := uuid.New().String()
	c.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    stateId,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		MaxAge:   600,
	})

	q := url.Values{}

	q.Set("client_id", clientId)
	q.Set("redirect_uri", "http://localhost:3000/api/user/google/response") // redirect url
	q.Set("response_type", "code")
	q.Set("scope", "email openid profile")
	q.Set("state", stateId)

	loginUrl := "https://accounts.google.com/o/oauth2/v2/auth?" + q.Encode()

	if err := c.Redirect().To(loginUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(commons.Results{
			Status: false,
			Msg:    "잘못된 요청입니다.",
		})
	}
	return nil
}

func (o OauthHandler) LoginWithGoogleResponse(c fiber.Ctx) error {
	ctx := context.Background()
	state := c.Query("state")
	code := c.Query("code")

	cookieState := c.Cookies("oauth_state")

	if state != cookieState {
		return c.Status(fiber.StatusUnauthorized).JSON(commons.Results{
			Status: false,
			Msg:    "같은 요청이 아닙니다.",
		})
	}
	c.ClearCookie("oauth_state")

	if code == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(commons.Results{
			Status: false,
			Msg:    "로그인 인증에 문제가 생겼습니다.",
		})
	}

	q := url.Values{}
	q.Set("client_id", os.Getenv("GOOGLE_CLIENT_ID"))
	q.Set("client_secret", os.Getenv("GOOGLE_SECRET_KEY"))
	q.Set("code", code)
	q.Set("grant_type", "authorization_code")
	q.Set("redirect_uri", "http://localhost:3000/api/user/google/response")

	resp, err := http.PostForm("https://oauth2.googleapis.com/token", q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(commons.Results{
			Status: false,
			Msg:    "토큰 교환에 실패했습니다.",
		})
	}
	defer resp.Body.Close()

	var tokenResp googleTokenResponse

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(commons.Results{
			Status: false,
			Msg:    "토큰 파싱에 실패했습니다.",
		})
	}

	token, _, err := jwt.NewParser().ParseUnverified(tokenResp.IdToken, &googleClaims{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(commons.Results{
			Status: false,
			Msg:    "토큰 파싱에 실패했습니다.",
		})
	}
	claims := token.Claims.(*googleClaims)

	email := claims.Email
	name := claims.Name

	exist, err := o.userService.FindUserByEmail(email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(commons.Results{
			Status: false,
			Msg:    "유저를 확인 하는 중 에러가 발생했습니다.",
		})
	}

	if exist == email {
		// JWT 등록
		commons.RegistCookies(c, email, name, "GOOGLE")
		return c.Status(fiber.StatusOK).JSON(commons.Results{
			Status: true,
			Msg:    "로그인에 성공하였습니다.",
			Data:   nil,
		})
	}

	id, err := uuid.NewV7()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var newUser request.CreateUserRequest

	newUser.Id = id.String()
	newUser.Email = email
	newUser.Social = "Google"
	newUser.Nickname = name

	if err := o.userService.SCreateUser(ctx, newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(commons.Results{
			Status: false,
			Msg:    "유저 생성에 실패하였습니다.",
		})
	}

	// JWT 등록
	commons.RegistCookies(c, email, name, "GOOGLE")

	return c.Status(fiber.StatusOK).JSON(commons.Results{
		Status: true,
		Msg:    "로그인에 성공하였습니다.",
		Data:   nil,
	})

}
