package handler

import (
	"fmt"
	"hexapi/databases"
	"hexapi/eservice"
	"hexapi/logs"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func ErrorCheck(err error) *fiber.Error {
	if err == databases.ErrDB {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "database error cannot get data")
	} else if err == databases.ErrNoRows {
		return fiber.NewError(fiber.StatusNotFound, "no row in db")
	} else if err == eservice.ErrNoDATAINPUT {
		return fiber.NewError(fiber.StatusBadRequest, "no id in request")
	} else if err == eservice.ErrNoDataForUpdate {
		return fiber.NewError(fiber.StatusBadRequest, "no data for update in request")
	} else if err == eservice.ErrProcessInterrup {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else {
		return fiber.NewError(fiber.StatusLengthRequired, "unpexted error")
	}
}

type MyCustomClaims struct {
	Id    int
	Rigth []string
	jwt.RegisteredClaims
}

func NewTokenKey(id int, rigth []string) (string, error) {
	key := []byte(viper.GetString("private_key"))
	claims := MyCustomClaims{
		id,
		rigth,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		},
	}
	ready := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := ready.SignedString(key)
	if err != nil {
		logs.Log.Error(err.Error())
		return "", err
	}
	return token, nil
}

func HaveAuthorization(c *fiber.Ctx) error {
	bearer := c.GetReqHeaders()["Authorization"]
	if len(bearer) == 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "please login ! and send token")
	}
	tokenReq := strings.Split(bearer[0], " ")
	tokenString := tokenReq[1]

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("private_key")), nil
	})
	if err != nil {
		logs.Log.Error(err.Error())
		return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("please send token\n%v", err.Error()))
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		c.Locals("id", claims.Id)
		c.Locals("rigth", claims.Rigth)
		return c.Next()
	} else {
		return fiber.ErrUnprocessableEntity
	}
}
