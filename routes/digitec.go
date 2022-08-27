package routes

import (
	"github.com/gaelgoth/swiss-deals-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func DigitecRoute(route fiber.Router) {
	route.Get("/digitec", controllers.GetDigitecOffer)
	route.Get("/galaxus", controllers.GetGalaxusOffer)
}
