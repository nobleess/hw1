package http

import (
	"context"
	"main/internal/user/domain/usecase"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type UserAPI struct {
	log     zerolog.Logger
	useCase usecase.Service
	cfg     *config.Notification
}

func NewUserAPI(cfg *config.Notification, router *mux.Router, useCase usecase.Service) *UserAPI {

	logger := log.With().
		Str("module", "notification-server").
		Logger()

	router.HandleFunc("/api/v1/users", useCase.Func).Methods("POST")
	return &UserAPI{
		log:     logger,
		useCase: useCase,
		cfg:     cfg,
	}
}
