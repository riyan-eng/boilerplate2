package middleware

import (
	"boilerplate/pkg"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func AuthorizeCasbin(enforce *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get current user
		userID, ok := c.Locals("userID").(string)
		if userID == "" || !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg.ResponseJson{
				Data:    "current logged in user not found",
				Message: pkg.MESSAGE_UNAUTHORIZED,
			})
		}

		// load policy
		if err := enforce.LoadPolicy(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(pkg.ResponseJson{
				Data:    "failed to load casbin policy.",
				Message: pkg.MESSAGE_BAD_SYSTEM,
			})
		}

		// casbin enforce policy
		accepted, err := enforce.Enforce(userID, c.OriginalURL(), c.Method()) // userID - url - method
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(pkg.ResponseJson{
				Data:    "error when authorizing user's accessibility.",
				Message: pkg.MESSAGE_BAD_SYSTEM,
			})
		}
		if !accepted {
			return c.Status(fiber.StatusForbidden).JSON(pkg.ResponseJson{
				Message: "you're not allowed.",
			})
		}
		return c.Next()
	}
}
