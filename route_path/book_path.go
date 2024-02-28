package route_path

import (
	"hexapi/handler"

	"github.com/gofiber/fiber/v2"
)

func BookRoute(app *fiber.App, h handler.BookHandler) {
	app.Post("/books", h.AddBook)
	app.Get("/books", h.GetAllBook)
	app.Get("/books/:id", h.GetByIdBook)
	app.Put("/books/:id", h.EditBookById)
}
