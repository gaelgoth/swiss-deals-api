package main

import (
	"log"
	"os"
	"time"

	_ "github.com/gaelgoth/swiss-deals-api/docs"
	"github.com/gaelgoth/swiss-deals-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/redirect/v2"
	"github.com/gofiber/swagger" // fiber-swagger middleware
	"github.com/joho/godotenv"
)

// @title       Swiss Deals API
// @version     1.0.0
// @description Aggregate deals of the day from Digitec, Galaxus, QoQa. Front-end available on http://deals.gothuey.dev/

// @contact.name Gaël G.
// @contact.url  https://blog.gothuey.dev/

// @license.name MIT license
// @license.url  https://github.com/gaelgoth/swiss-deals-api/blob/main/LICENSE

// @host deals-api.gothuey.dev
// //@host localhost:9000

func main() {
	err := godotenv.Load()

	port := os.Getenv("PORT")
	app := fiber.New()

	app.Use(
		logger.New(), cors.New(),
		redirect.New(redirect.Config{
			Rules: map[string]string{
				"/": "/swagger/index.html",
			},
			StatusCode: 301,
		}),
		cache.New(cache.Config{
			// caching requests during 15 minutes
			Expiration:   15 * time.Minute,
			CacheControl: true,
		}),
	)
	setupRoutes(app)
	app.Get("/swagger/*", swagger.HandlerDefault)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Fatal(app.Listen(":" + port))

}

// HealthCheck godoc
// @Summary     Show the status of server.
// @Description get the status of server.
// @Tags        Health check 🩺
// @Accept      */*
// @Produce     json
// @Router      / [get]
func setupRoutes(app *fiber.App) {

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Everything Is Okay 👋🏾 from air!",
		})
	})

	// api group
	api := app.Group("/api")

	routes.DigitecRoute(api.Group("/deals"))
	routes.QoqaRoute(api.Group("/deals"))
}
