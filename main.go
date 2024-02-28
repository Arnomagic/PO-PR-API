package main

import (
	"fmt"
	"hexapi/configuration"
	"hexapi/databases"
	"hexapi/eservice"
	"hexapi/handler"
	"hexapi/route_path"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	configuration.InitViperConfig()
	db := configuration.InitDb()
	if db == nil {
		return
	}
	configuration.InitTimeZone()
	app := fiber.New()

	authDB := databases.NewAuthDb(db)
	authEservice := eservice.NewAuthEservice(authDB)
	authHandler := handler.NewAuthHandler(authEservice)
	app.Post("/login", authHandler.Login)

	memberDB := databases.NewMemberDatabases(db)
	memberEservice := eservice.NewMemberEservice(memberDB)
	memberHandler := handler.NewMemberHandler(memberEservice)
	route_path.MemberRoute(app, memberHandler)

	app.Listen(":" + fmt.Sprint(viper.GetInt("app.port")))

}
