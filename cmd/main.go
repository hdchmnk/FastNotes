package main

import (
	_ "FastNotes/cmd/docs"
	"FastNotes/internal/db"
	"FastNotes/internal/user"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
)

// @title FastNotes Swagger API
// @version 1.0
// @description This is Swagger API for FastNotes API
// @host localhost:8080
// @BasePath /

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("db connection error")
	}

	e := echo.New()

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	e.POST("/signup", userHandler.CreateUser)
	e.POST("/login", userHandler.Login)
	e.GET("/logout", userHandler.Logout)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server error")
	}
}
