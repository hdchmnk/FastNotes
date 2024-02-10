package main

import (
	_ "FastNotes/docs"
	"FastNotes/internal/db"
	"FastNotes/internal/notes"
	"FastNotes/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

// @title FastNotes Swagger API
// @version 1.0
// @description This is Swagger for FastNotes API
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

	notesRepository := notes.NewRepository(dbConn.GetDB())
	notesService := notes.NewService(notesRepository)
	notesHandler := notes.NewHandler(notesService)

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	r.POST("/createnote", notesHandler.CreateNote)
	r.POST("/getnotesbyid", notesHandler.GetNotesByUserID)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server error")
	}
}
