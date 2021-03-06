package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	f "github.com/Learning-Management-System-Kelompok-42/BE-LMS/helpers/formatter"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type extract struct {
	CompanyId   string
	UserId      string
	LevelAccess string
}

var jwtSignedMethod = jwt.SigningMethodHS256

func JWTMiddleware(config *config.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// if strings.Contains(c.Request().URL.Path, "/login") {
			// 	return next(c)
			// }

			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, f.ForbiddenResponse("Invalid token"))
			}
			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, f.ForbiddenResponse("Invalid token"))
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.App.SecretKey), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || !token.Valid || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, f.ForbiddenResponse("Invalid Token"))
			}

			return next(c)
		}
	}
}

func ExtractToken(c echo.Context) (result extract, err error) {

	signature := strings.Split(c.Request().Header.Get("Authorization"), " ")

	claim := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
		return []byte("Secret_JWT"), nil
	})

	result.CompanyId = fmt.Sprintf("%v", claim["CompanyID"])
	result.UserId = fmt.Sprintf("%v", claim["UserID"])
	result.LevelAccess = fmt.Sprintf("%v", claim["LevelAccess"])

	return result, nil
}

// Adding handlerFunction to check if the LevelAccess == admin or user
func CheckLevelAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		extract, err := ExtractToken(c)

		if err != nil {
			return c.JSON(http.StatusForbidden, f.ForbiddenResponse("Invalid token"))
		}

		if extract.LevelAccess != "admin" {
			return c.JSON(http.StatusUnauthorized, f.UnauthorizedResponse("You are not authorized to access this route"))
		}

		return next(c)
	}
}
