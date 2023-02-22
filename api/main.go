package main

import (
	"clean-architecture-golang/internal/config"
	"clean-architecture-golang/internal/handler"
	"clean-architecture-golang/internal/infrastructure"
	"clean-architecture-golang/internal/middleware"
	"clean-architecture-golang/internal/repository"
	"clean-architecture-golang/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	appConfig := config.NewAppConfig()
	db := infrastructure.NewDatabaseConnection()

	userRepo := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)

	r := gin.Default()
	r.Use(middleware.LogErrors())

	handler.MakeUserHandler(r, userService)

	if err := r.Run(":" + appConfig.Port); err != nil {
		log.Fatal().Err(err).Send()
	}
}
