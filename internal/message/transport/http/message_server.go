package http

import (
	"main/internal/user/domain/usecase"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type MessageAPI struct {
	log     zerolog.Logger
	useCase usecase.UseCase
}

func NewMessageAPI(router *mux.Router, useCase usecase.UseCase) {

	logger := log.With().
		Str("module", "notification-server").
		Logger()

	MessageAPI := MessageAPI{
		log:     logger,
		useCase: useCase,
	}
	router.HandleFunc("/api/v1/message", MessageAPI.Create).Methods("POST")
	router.HandleFunc("/api/v1/message", MessageAPI.Read).Methods("GET")

	router.HandleFunc("/api/v1/message/{id}", MessageAPI.Update).Methods("PUT")
	router.HandleFunc("/api/v1/message/{id}", MessageAPI.Delete).Methods("DELETE")
	router.HandleFunc("/api/v1/message/{id}", MessageAPI.Read).Methods("POST")
}
