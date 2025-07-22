package api

import (
	"context"
	"net/http"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/LouisFernando1204/golang-restapi.git/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type journalApi struct {
	journalService domain.JournalService
}

func NewJournal(app *fiber.App, journalService domain.JournalService, authMiddleware fiber.Handler) {

	ja := journalApi{
		journalService: journalService,
	}

	journal := app.Group("/journals", authMiddleware)
	journal.Get("", ja.Index)
	journal.Post("", ja.Create)
	journal.Put(":id", ja.Update)
}

func (ja journalApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	customerId := ctx.Query("customer_id")
	status := ctx.Query("status")

	res, err := ja.journalService.Index(c, domain.JournalSearch{
		CustomerId: customerId,
		Status:     status,
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusOK
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, res))
}

func (ja journalApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateJournalRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		statusCode := http.StatusBadRequest
		return ctx.Status(statusCode).JSON(dto.CreateResponseErrorData(statusCode, "Validation failed", fails))
	}

	err := ja.journalService.Create(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusCreated
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, ""))
}

func (ja journalApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	userToken := ctx.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	err := ja.journalService.Return(c, dto.ReturnJournalRequest{
		JournalId: id,
		UserId:    userId,
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusOK
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, ""))
}
