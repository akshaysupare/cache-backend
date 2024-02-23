package api

import (
	"fmt"
	"strconv"
	"time"

	"backend-config.Cache/config"

	"github.com/gofiber/fiber/v2"
)

func SetKey(c *fiber.Ctx) error {
	key := c.Query("key")
	val := c.Query("val")
	expiry := c.Query("expiry")

	// var content Content
	// if err := c.ShouldBindJSON(&content); err != nil {
	// 	c.Error(err)
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }
	if key != "" && val != "" && expiry != "" {
		msg := ""
		exp, err := strconv.Atoi(expiry)
		if err != nil {
			fmt.Println("failed to get expiry, setting up default expiry of 5 Sec")
			exp = 5
			msg = " Unable to get expiry; defaulting to 5 sec"
		}

		ok := config.Cache.SetWithTTL(key, val, 1, time.Duration(exp)*time.Second)
		if ok {
			c.Status(200).JSON(
				&fiber.Map{
					"message": "success " + msg,
				},
			)

		} else {
			c.Status(500).JSON(
				&fiber.Map{
					"message": "failed to add key in cache",
				},
			)
		}
	}else  {
		c.Status(400).JSON(
			&fiber.Map{
				"message": "Bad Request!",
			},
		)
	}
	return nil
}
