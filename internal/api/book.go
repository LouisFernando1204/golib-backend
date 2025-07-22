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

type bookApi struct {
	bookService domain.BookService
}

func NewBook(app *fiber.App, bookService domain.BookService, authMiddleware fiber.Handler) {
	ba := bookApi{
		bookService: bookService,
	}

	book := app.Group("/books", authMiddleware)
	book.Get("", ba.Index)
	book.Get(":id", ba.Show)
	book.Post("", ba.Create)
	book.Put(":id", ba.Update)
	book.Delete(":id", ba.Delete)

	// app.Get("/books", authMiddleware, ba.Index)
	// app.Get("/books/:id", authMiddleware, ba.Show)
	// app.Post("/books", authMiddleware, ba.Create)
	// app.Put("/books/:id", authMiddleware, ba.Update)
	// app.Delete("/books/:id", authMiddleware, ba.Delete)
}

func (ba bookApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ba.bookService.Index(c)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusOK
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, res))
}

func (ba bookApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	data, err := ba.bookService.Show(c, id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusOK
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, data))
}

func (ba bookApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.CreateResponseErrorData(statusCode, "Validation failed", fails))
	}

	err := ba.bookService.Create(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusCreated
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, ""))
}

func (ba bookApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.CreateResponseErrorData(statusCode, "Validation failed", fails))
	}

	req.Id = ctx.Params("id")
	err := ba.bookService.Update(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusOK
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, ""))
}

func (ba bookApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ba.bookService.Delete(c, id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.SendStatus(http.StatusNoContent)
}
