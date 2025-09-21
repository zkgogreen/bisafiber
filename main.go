package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zkgogreen/bisago/internal/api"
	"github.com/zkgogreen/bisago/internal/config"
	"github.com/zkgogreen/bisago/internal/connection"
	"github.com/zkgogreen/bisago/internal/repository"
	"github.com/zkgogreen/bisago/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()
	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)
	api.NewCustomer(app, customerService)
	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
