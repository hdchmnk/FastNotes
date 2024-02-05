package main

import (
	_ "FastNotes/cmd/docs"
	"FastNotes/internal/db"
	"FastNotes/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	r := gin.New()

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	if err := r.Run(":8080"); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server error")
	}
}
