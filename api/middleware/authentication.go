package middleware

import (
	"go-covid/business/admin"
	"go-covid/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupAuthenticationJWT() echo.MiddlewareFunc {
	SECRET_KEY := config.GetConfig().Secrettoken.Token
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		Claims:        &admin.Claims{},
		SigningKey:    []byte(SECRET_KEY),
	})
}
