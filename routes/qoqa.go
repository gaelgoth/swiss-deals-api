package routes

import (
	"github.com/gael-gothuey/deal-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func QoqaRoute(route fiber.Router) {
	route.Get("/qoqa", controllers.GetQoqaOffer)
	route.Get("/qwine", controllers.GetQwineOffer)
	route.Get("/qbeer", controllers.GetBeerOffer)
	route.Get("/qsport", controllers.GetQsportOffer)
	route.Get("/qooking", controllers.GetQookingOffer)
	route.Get("/qlock", controllers.GetQlockOffer)
}
