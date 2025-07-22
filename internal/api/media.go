package api

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/LouisFernando1204/golang-restapi.git/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type mediaApi struct {
	cnf          *config.Config
	mediaService domain.MediaService
}

func NewMedia(app *fiber.App, cnf *config.Config, mediaService domain.MediaService, authMiddleware fiber.Handler) {
	ma := mediaApi{
		cnf:          cnf,
		mediaService: mediaService,
	}

	app.Post("/media", authMiddleware, ma.Create)
	app.Static("/media", cnf.Storage.BasePath)
}

func (ma mediaApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	file, err := ctx.FormFile("media")
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	filename := uuid.NewString() + "_" + file.Filename
	path := filepath.Join(ma.cnf.Storage.BasePath, filename)
	err = ctx.SaveFile(file, path)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	res, err := ma.mediaService.Create(c, dto.CreateMediaRequest{
		Path: filename,
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	statusCode := http.StatusCreated
	return ctx.Status(statusCode).JSON(dto.CreateResponseSuccess(statusCode, res))
}
