package router

import (
	"log"

	"backend-config.Cache/api"
	"backend-config.Cache/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitRoutes() {

	router := fiber.New()

	//setting up cors policy
	router.Use(cors.New(cors.Config{
		AllowOrigins:  "*", // Update this with your allowed origins
		AllowMethods:  "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders: "Content-Length",
	}))

	router.Get("/get_key", api.GetKey)
	router.Post("/set_key", api.SetKey)

	//staring server 
	log.Fatal(router.Listen(":" + config.Port))

}
