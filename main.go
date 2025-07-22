package main

import (
	"errors"
	"net/http"

	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/LouisFernando1204/golang-restapi.git/internal/api"
	"github.com/LouisFernando1204/golang-restapi.git/internal/config"
	"github.com/LouisFernando1204/golang-restapi.git/internal/connection"
	"github.com/LouisFernando1204/golang-restapi.git/internal/repository"
	"github.com/LouisFernando1204/golang-restapi.git/internal/service"
	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	cnf, _ := config.Get()

	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	jwtMiddleware := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			message := "Unauthorized: Invalid or Malformed Token"
			if errors.Is(err, jwt.ErrTokenMalformed) {
				message = "Unauthorized: Token format is invalid" // Format token salah
			} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
				message = "Unauthorized: Token has expired" // Token kedaluwarsa
			} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
				message = "Unauthorized: Token signature is invalid" // Tanda tangan tidak cocok (palsu)
			} else if err.Error() == "missing or malformed JWT" {
				message = "Unauthorized: Missing or malformed Bearer token" // Tidak ada token
			}
			statusCode := http.StatusUnauthorized
			return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, message))
		},
	})

	customerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)
	bookRepository := repository.NewBook(dbConnection)
	bookStockRepository := repository.NewBookStock(dbConnection)
	journalRepository := repository.NewJournal(dbConnection)
	mediaRepository := repository.NewMedia(dbConnection)
	chargeRepository := repository.NewCharge(dbConnection)

	authService := service.NewAuth(cnf, userRepository)
	customerService := service.NewCustomer(customerRepository)
	bookService := service.NewBook(cnf, bookRepository, bookStockRepository, mediaRepository)
	bookStockService := service.NewBookStock(bookRepository, bookStockRepository)
	journalService := service.NewJournal(journalRepository, bookRepository, bookStockRepository, customerRepository, chargeRepository)
	mediaService := service.NewMedia(cnf, mediaRepository)

	api.NewAuth(app, authService)
	api.NewCustomer(app, customerService, jwtMiddleware)
	api.NewBook(app, bookService, jwtMiddleware)
	api.NewBookStock(app, bookStockService, jwtMiddleware)
	api.NewJournal(app, journalService, jwtMiddleware)
	api.NewMedia(app, cnf, mediaService, jwtMiddleware)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
