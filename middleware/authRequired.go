package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/hecharles/doctortest/config"
)

//AuthRequired handler
func AuthRequired() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.JwtSecret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")

		},
	})
}
