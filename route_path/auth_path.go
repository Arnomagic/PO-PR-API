package route_path

import (
	"hexapi/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, h handler.AuthHandler) {
	app.Patch("/rigths", h.AddRigthToId)
	app.Get("/rigths/:id", h.GetRigthById)
	app.Delete("/rigths/:id/:index", h.RemoveRigthByIdFromIndex)
}
