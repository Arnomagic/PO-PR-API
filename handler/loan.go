package handler

import (
	"hexapi/databases"
	"hexapi/eservice"

	"github.com/gofiber/fiber/v2"
)

type LoanHandler interface {
	AddLoan(c *fiber.Ctx) error
	GetAllLoan(c *fiber.Ctx) error
	GetByIdLoan(c *fiber.Ctx) error
	EditLoanyId(c *fiber.Ctx) error
}
type loanHandler struct {
	d eservice.LoanEservice
}

func NewLoanHandler(s eservice.LoanEservice) LoanHandler {
	return loanHandler{d: s}
}

func (h loanHandler) AddLoan(c *fiber.Ctx) error {
	loan := eservice.LoanInsert{}
	content := c.GetReqHeaders()[fiber.HeaderContentType]
	if content[0] != fiber.MIMEApplicationJSON {
		return fiber.NewError(fiber.StatusBadRequest, "please send json. header")
	}
	err := c.BodyParser(&loan)
	if err != nil {
		return fiber.ErrBadRequest
	}
	res, err := h.d.AddLoan(loan)
	if err != nil {
		return ErrorCheck(err)
	}
	c.Status(fiber.StatusCreated)
	return c.JSON(res)
}
func (h loanHandler) GetAllLoan(c *fiber.Ctx) error {
	response, err := h.d.GetAllLoan()
	if err != nil {
		return ErrorCheck(err)
	}
	return c.JSON(response)
}
func (h loanHandler) GetByIdLoan(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is integer please check url")
	}
	response, err := h.d.GetByIdLoan(id)
	if err != nil {
		return ErrorLoan(err)
	}
	return c.JSON(response)
}
func (h loanHandler) EditLoanyId(c *fiber.Ctx) error {
	loan := eservice.Loan{}
	content := c.GetReqHeaders()[fiber.HeaderContentType]
	if content[0] != fiber.MIMEApplicationJSON {
		return fiber.NewError(fiber.StatusBadRequest, "please send json. header")
	}
	err := c.BodyParser(&loan)
	if err != nil {
		return fiber.ErrBadRequest
	}
	res, err := h.d.EditLoanById(loan)
	if err != nil {
		return ErrorLoan(err)
	}
	return c.JSON(res)
}
func ErrorLoan(err error) *fiber.Error {
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
