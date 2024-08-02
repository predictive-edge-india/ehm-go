package middlewares

import (
	"github.com/predictive-edge-india/ehm-go/helpers"
	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

const SECRET = "1ZVtHIAjEPcNrSHQ3lyWdqxFlsSA81YhdbXK7D57c0L19kxq27CqIQvkzxkPcMj6YARFVuvE8PLOeTGZbbjoWU904vdSCOA2dwaMd05cmnWt3iS4Xk0w9z1kFLtzcF"

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(SECRET),
		},
		ErrorHandler: jwtError,
		AuthScheme:   "JWT",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	log.Error().AnErr("jwtError:", err).Send()
	if err.Error() == "Missing or malformed JWT" {
		return helpers.NotAuthenticatedError(c, "Missing or malformed JWT")
	} else {
		return helpers.NotAuthenticatedError(c, "Invalid or expired JWT")
	}
}
