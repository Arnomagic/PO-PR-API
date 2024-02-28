package handler

import (
	"hexapi/databases"
	"hexapi/eservice"
	"hexapi/logs"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	AddRigthToId(c *fiber.Ctx) error
	GetRigthById(c *fiber.Ctx) error
	RemoveRigthByIdFromIndex(c *fiber.Ctx) error
}
type authHandler struct {
	e eservice.AuthEservice
}

func NewAuthHandler(e eservice.AuthEservice) AuthHandler {
	return authHandler{e: e}
}

func (h authHandler) Login(c *fiber.Ctx) error {
	userData := eservice.Credential{}
	err := c.BodyParser(&userData)
	if err != nil {
		logs.Log.Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "please send data correct !")
	}
	res, err := h.e.GetUserByCredentail(userData)
	if err != nil {
		if err == databases.ErrNoRows {
			return fiber.NewError(fiber.StatusUnauthorized, "email or password invalid!")
		}
		return ErrorCheck(err)
	}
	token, err := NewTokenKey(res.Id, res.Rigth)
	if err != nil {
		logs.Log.Error(err.Error())
		return ErrorCheck(eservice.ErrProcessInterrup)
	}
	return c.JSON(fiber.Map{
		"message": "login sucessfully!",
		"id":      res.Id,
		"token":   token,
	})
}
func (h authHandler) AddRigthToId(c *fiber.Ctx) error {
	if !isRigth(c.Locals("rigth").([]string), "addRigth") {
		return fiber.NewError(fiber.StatusUnauthorized, "no rigth to add data!")
	}
	rigth := eservice.CredentialResponse{}
	err := c.BodyParser(&rigth)
	if err != nil {
		logs.Log.Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "please send corect")
	}
	res, err := h.e.AddRigthToId(rigth)
	if err != nil {
		return nil
	}
	return c.JSON(fiber.Map{
		"message": "sucessfully!",
		"rigth":   res,
	})
}
func (h authHandler) GetRigthById(c *fiber.Ctx) error {
	if !isRigth(c.Locals("rigth").([]string), "showRigth") {
		return fiber.NewError(fiber.StatusUnauthorized, "no rigth to add showRigth!")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.Log.Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest)
	}
	res, err := h.e.GetRigthByid(id)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil
	}
	return c.JSON(fiber.Map{
		"message": "sucessfully!",
		"rigth":   res,
	})
}
func (h authHandler) RemoveRigthByIdFromIndex(c *fiber.Ctx) error {
	if !isRigth(c.Locals("rigth").([]string), "deleteRigth") {
		return fiber.NewError(fiber.StatusUnauthorized, "no rigth to add deleteRigth!")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	index, err := c.ParamsInt("index")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	err = h.e.RemoveRigthByIdFromIndex(id, index)
	if err != nil {
		return ErrorCheck(err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
