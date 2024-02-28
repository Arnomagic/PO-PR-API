package route_path

import (
	"hexapi/handler"

	"github.com/gofiber/fiber/v2"
)

func MemberRoute(app *fiber.App, h handler.MemberHandler) {
	app.Use(handler.HaveAuthorization)
	app.Post("/members", h.AddMember)
	app.Get("/members", h.GetAllMember)
	app.Get("/members/:id", h.GetByIdMember)
	app.Put("/members", h.EditMemberById)
}
