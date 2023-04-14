package middleware

import (
	"boilerplate/config"
	"boilerplate/internal/util"
	"boilerplate/pkg"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/errgroup"
)

func AuthorizeJwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
				Data:    "authorization header is required",
				Message: "bad",
			})
		}
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 {
			return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
				Data:    "undefined token",
				Message: "bad",
			})
		}
		tokenString := splitToken[1]
		claims, err := util.ParseToken(tokenString, config.GetEnv("JWT_SECRET_ACCESS"))
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// Malformed token -> Delete Cookie
				// c.ClearCookie("jwt")
				return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
					Data:    "missing or malformed jwt",
					Message: "ok",
				})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return c.Status(fiber.StatusUnauthorized).JSON(pkg.ResponseJson{
					Data:    "expired token",
					Message: "unauthorized",
				})
			} else {
				// Cannot handle -> Delete Cookie
				c.ClearCookie("jwt")
				return c.Status(fiber.StatusForbidden).JSON(pkg.ResponseJson{
					Data:    "error when processing identity",
					Message: "forbidden",
				})
			}
		}

		if err := util.ValidateToken(claims, false, c.Context()); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg.ResponseJson{
				Data:    "token not authorized",
				Message: "unauthorized",
			})
		}
		c.Locals("userID", claims.UserID)
		c.Locals("companyID", claims.CompanyID)

		g := new(errgroup.Group)
		g.Go(func() (err error) {
			err = config.Redis.Expire(c.Context(), fmt.Sprintf("token-%s", claims.UserID), time.Minute*time.Duration(15)).Err()
			return
		})
		g.Wait()

		return c.Next()
	}
}
