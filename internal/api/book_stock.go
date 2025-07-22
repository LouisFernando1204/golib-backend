package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/LouisFernando1204/golang-restapi.git/internal/util"
	"github.com/gofiber/fiber/v2"
)

type bookStockApi struct {
	bookStockService domain.BookStockService
}

func NewBookStock(app *fiber.App, bookStockService domain.BookStockService, authMiddleware fiber.Handler) {
	bsa := bookStockApi{
		bookStockService: bookStockService,
	}

	bookStock := app.Group("/book-stocks", authMiddleware)
	bookStock.Post("", bsa.Create)
	bookStock.Delete("", bsa.Delete)
}

func (bsa bookStockApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookStockRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.CreateResponseErrorData(statusCode, "Validation failed", fails))
	}

	err := bsa.bookStockService.Create(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusCreated
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, ""))
}

func (bsa bookStockApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	codeStr := ctx.Query("codes")
	if codeStr == "" {
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, "Code parameter is required"))
	}
	codes := strings.Split(codeStr, ";")

	var req dto.DeleteBookStockRequest
	req.Codes = codes

	err := bsa.bookStockService.Delete(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.SendStatus(http.StatusNoContent)
}
