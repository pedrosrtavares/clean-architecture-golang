package main

import (
	"clean-architecture-golang/internal/config"
	"clean-architecture-golang/internal/handler"
	"clean-architecture-golang/internal/infrastructure"
	"clean-architecture-golang/internal/metric"
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
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	r := gin.Default()
	r.Use(middleware.LogErrors())
	r.Use(middleware.Metrics(metricService))

	handler.MakeUserHandler(r, userService)

	err = r.Run(":" + appConfig.Port)
	if err != nil {
	}
}
