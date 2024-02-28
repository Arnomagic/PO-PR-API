package handler

import (
	"hexapi/databases"
	"hexapi/eservice"

	"github.com/gofiber/fiber/v2"
)

type BookHandler interface {
	AddBook(c *fiber.Ctx) error
	GetAllBook(c *fiber.Ctx) error
	GetByIdBook(c *fiber.Ctx) error
	EditBookById(c *fiber.Ctx) error
}
type bookHandler struct {
	d eservice.BookEservice
}

func NewBookHandler(s eservice.BookEservice) BookHandler {
	return bookHandler{d: s}
}
func (h bookHandler) AddBook(c *fiber.Ctx) error {
	book := eservice.BookInsert{}
	content := c.GetReqHeaders()[fiber.HeaderContentType]
	if content[0] != fiber.MIMEApplicationJSON {
		return fiber.NewError(fiber.StatusBadRequest, "please send json. header")
	}
	err := c.BodyParser(&book)
	if err != nil {
		return fiber.ErrBadRequest
	}
	res, err := h.d.AddBook(book)
	if err != nil {
		return ErrorCheck(err)
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(res)
}
func (h bookHandler) GetAllBook(c *fiber.Ctx) error {
	response, err := h.d.GetAllBook()
	if err != nil {
		return ErrorCheck(err)
	}
	return c.JSON(response)
}
func (h bookHandler) GetByIdBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is integer please check url")
	}
	response, err := h.d.GetByIdBook(id)
	if err != nil {
		return ErrorBook(err)
	}
	return c.JSON(response)
}
func (h bookHandler) EditBookById(c *fiber.Ctx) error {
	book := eservice.Book{}
	content := c.GetReqHeaders()[fiber.HeaderContentType]
	if content[0] != fiber.MIMEApplicationJSON {
		return fiber.NewError(fiber.StatusBadRequest, "please send json. header")
	}
	err := c.BodyParser(&book)
	if err != nil {
		return fiber.ErrBadRequest
	}
	res, err := h.d.EditBookById(book)
	if err != nil {
		return ErrorBook(err)
	}
	return c.JSON(res)
}
func ErrorBook(err error) *fiber.Error {
	if err == databases.ErrDB {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "database error cannot get data")
	} else if err == databases.ErrNoRows {
		return fiber.NewError(fiber.StatusNotFound, "no row in db")
	} else if err == eservice.ErrNoDATAINPUT {
		return fiber.NewError(fiber.StatusBadRequest, "no id in request")
	} else if err == eservice.ErrNoDataForUpdate {
		return fiber.NewError(fiber.StatusBadRequest, "no data for update in request")
	} else {
		return fiber.NewError(fiber.StatusLengthRequired, "unpexted error")
	}
}
