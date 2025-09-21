package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zkgogreen/bisago/domain"
	"github.com/zkgogreen/bisago/dto"
	"github.com/zkgogreen/bisago/internal/util"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := &customerApi{
		customerService: customerService,
	}
	app.Get("/customer", ca.Index)
	app.Post("/customer", ca.Create)
}

func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cencel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cencel()
	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}

func (ca customerApi) Create(ctx *fiber.Ctx) error {
	c, cencel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cencel()
	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnavailableForLegalReasons)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("error", fails))
	}
	err := ca.customerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(dto.CustomerData{
		Code: req.Code,
		Name: req.Name,
	}))
}
