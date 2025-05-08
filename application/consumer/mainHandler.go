package consumer

import (
	"com.github/confusionhill-aqw-ps/internal/config"
	"com.github/confusionhill-aqw-ps/internal/domain/auth"
)

type Handlers struct {
	Auth *auth.Handler
}

func NewHandlers(cfg *config.Config, usecase *Usecases) (*Handlers, error) {
	authHandler, err := auth.NewHandler(cfg, usecase.auth)
	if err != nil {
		return nil, err
	}
	return &Handlers{Auth: authHandler}, nil
}
