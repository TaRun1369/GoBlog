package middleware

import (
	"github.com/TaRun1369/GoBlog/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if _, err := util.Parsejwt(cookie); err != nil {
		// here we are checking if the cookie is valid or not
		//by checking if the cookie is valid or not we are checking if the user is authenticated or not
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	return c.Next() // if the cookie is valid then we will call the next function
	// next function gives the control to the next function which routes to next router to run
}
