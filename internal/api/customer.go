package api

import (
	"context"
	"net/http"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/LouisFernando1204/golang-restapi.git/internal/util"
	"github.com/gofiber/fiber/v2"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService, authMiddleware fiber.Handler) {
	ca := customerApi{
		customerService: customerService,
	}

	customer := app.Group("/customers", authMiddleware)
	customer.Get("", ca.Index)
	customer.Get(":id", ca.Show)
	customer.Post("", ca.Create)
	customer.Put(":id", ca.Update)
	customer.Delete(":id", ca.Delete)
}

func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusOK
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, res))
}

func (ca customerApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	data, err := ca.customerService.Show(c, id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusOK
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, data))
}

func (ca customerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.CreateResponseErrorData(statusCode, "Validation failed", fails))
	}

	err := ca.customerService.Create(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusCreated
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, ""))
}

func (ca customerApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.CreateResponseErrorData(statusCode, "Validation failed", fails))
	}

	req.ID = ctx.Params("id")
	err := ca.customerService.Update(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusOK
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, ""))
}

func (ca customerApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ca.customerService.Delete(c, id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.SendStatus(http.StatusNoContent)
}
