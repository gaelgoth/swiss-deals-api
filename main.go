package main

import (
	"log"
	"os"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2" // fiber-swagger middleware
	"github.com/joho/godotenv"

	_ "github.com/gael-gothuey/deal-backend/docs"
	"github.com/gael-gothuey/deal-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/redirect/v2"
)

// @title        Swiss Deals API
// @version      1.0.0
// @description  This API aggregates promotions from different Swiss online store.

// @contact.name   Gaël G.
// @contact.url    https://blog.gothuey.dev/

// @license.name  MIT license
// @license.url   https://github.com/gaelgoth/swiss-deals-api/blob/main/LICENSE

// @host deals-api.gothuey.dev
// //@host localhost:9000

// @BasePath  /api

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
// @Summary      Show the status of server.
// @Description  get the status of server.
// @Tags         Health check 🩺
// @Accept       */*
// @Produce      json
// @Router       / [get]
func setupRoutes(app *fiber.App) {

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/api")

	routes.DigitecRoute(api.Group("/deals"))
	routes.QoqaRoute(api.Group("/deals"))
}
