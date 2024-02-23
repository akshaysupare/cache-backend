package api

import (
	"backend-config.Cache/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetKey(c *fiber.Ctx) error {
	key := c.Query("key")

	fmt.Println(" key = ", key)

	val, ok := config.Cache.Get(key)
	if ok {

		c.Status(200).JSON(&fiber.Map{
			"key ": key,
			"val":  val,
		})

	} else {
		c.Status(400).JSON(
			&fiber.Map{
				"message": "Bad Request!",
			},
		)
	}
	return nil
}
