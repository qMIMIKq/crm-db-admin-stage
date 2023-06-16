package main

import (
	crm "crm_admin"
	"crm_admin/internal/config"
	"crm_admin/internal/repository"
	"crm_admin/internal/services"
	"crm_admin/internal/transport"
	"crm_admin/pkg/database"
	"crm_admin/pkg/logger"
	"github.com/rs/zerolog/log"
)

func init() {
	logger.CustomizeLogger()
}

func main() {
	cfg := config.GetConfig()

	pgDb, err := database.NewPostgresDB(cfg.CrmDB)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("error connecting to database")
	}

	repos := repository.NewRepository(pgDb)
	service := services.NewService(repos)
	mainHandler := transport.NewMainHandler(service)

	srv := crm.NewServer(cfg.AppPort, mainHandler.InitAllRoutes())
	if err := srv.Run(); err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}
