package handler

import (
	"hexapi/eservice"

	"github.com/gofiber/fiber/v2"
)

type MemberHandler interface {
	AddMember(c *fiber.Ctx) error
	GetAllMember(c *fiber.Ctx) error
	GetByIdMember(c *fiber.Ctx) error
	EditMemberById(c *fiber.Ctx) error
}
type memberHandler struct {
	e eservice.MemberEservice
}

func NewMemberHandler(s eservice.MemberEservice) MemberHandler {
	return memberHandler{e: s}
}
func (h memberHandler) AddMember(c *fiber.Ctx) error {
	isRigth := isRigth(c.Locals("rigth").([]string), "add")
	if !isRigth {
		return fiber.NewError(fiber.StatusUnauthorized, "no rigth to add data!")
	}
	member := eservice.MemberInsert{}
	content := c.GetReqHeaders()[fiber.HeaderContentType]
	if content[0] != fiber.MIMEApplicationJSON {
		return fiber.NewError(fiber.StatusBadRequest, "please send json. header")
	}
	err := c.BodyParser(&member)
	if err != nil {
		return fiber.ErrBadRequest
	}
	res, err := h.e.AddMember(member)
	if err != nil {
		return ErrorCheck(err)
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(res)
}
func (h memberHandler) GetAllMember(c *fiber.Ctx) error {
	isRigth := isRigth(c.Locals("rigth").([]string), "show")
	if !isRigth {
		return fiber.NewError(fiber.StatusUnauthorized, "no rigth to show data!")
	}
	response, err := h.e.GetAllMember()
	if err != nil {
		return ErrorCheck(err)
	}
	return c.JSON(response)
}
func (h memberHandler) GetByIdMember(c *fiber.Ctx) error {
	isRigth := isRigth(c.Locals("rigth").([]string), "showById")
	if !isRigth {
		return fiber.NewError(fiber.StatusUnauthorized, "no rigth to show data by id!")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is integer please check url")
	}
	response, err := h.e.GetByIdMember(id)
	if err != nil {
		return ErrorCheck(err)
	}
	return c.JSON(response)
}
func (h memberHandler) EditMemberById(c *fiber.Ctx) error {
	member := eservice.Member{}
	content := c.GetReqHeaders()[fiber.HeaderContentType]
	if content[0] != fiber.MIMEApplicationJSON {
		return fiber.NewError(fiber.StatusBadRequest, "please send json. header")
	}
	err := c.BodyParser(&member)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if c.Locals("id").(int) == member.MemberID {

	} else {
		isRigth := isRigth(c.Locals("rigth").([]string), "edit")
		if !isRigth {
			return fiber.NewError(fiber.StatusUnauthorized, "no rigth to edit data!")
		}
	}

	res, err := h.e.EditMemberById(member)
	if err != nil {
		return ErrorCheck(err)
	}
	return c.JSON(res)
}
func isRigth(rigth []string, word string) bool {
	for _, row := range rigth {
		if row == word {
			return true
		}
	}
	return false
}
