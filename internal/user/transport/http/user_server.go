package http

import (
	"main/internal/user/domain/usecase"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type UserAPI struct {
	log     zerolog.Logger
	useCase usecase.UseCase
}

func NewUserAPI(router *mux.Router, useCase usecase.UseCase) {

	logger := log.With().
		Str("module", "notification-server").
		Logger()

	userAPI := UserAPI{
		log:     logger,
		useCase: useCase,
	}
	router.HandleFunc("/api/v1/users", userAPI.Create).Methods("POST")
	router.HandleFunc("/api/v1/users", userAPI.Read).Methods("GET")

	router.HandleFunc("/api/v1/users/{id}", userAPI.Update).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", userAPI.Delete).Methods("DELETE")
	router.HandleFunc("/api/v1/users/{id}", userAPI.Read).Methods("POST")
}
