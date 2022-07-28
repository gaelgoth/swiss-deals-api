package routes

import (
	"github.com/gael-gothuey/deal-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func DigitecRoute(route fiber.Router) {
	route.Get("/digitec", controllers.GetDigitecOffer)
	route.Get("/galaxus", controllers.GetGalaxusOffer)
}
