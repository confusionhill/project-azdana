package consumer

import (
	"com.github/confusionhill-aqw-ps/internal/config"
	"com.github/confusionhill-aqw-ps/internal/domain/auth"
)

type Usecases struct {
	auth *auth.Usecase
}

func NewUsecases(cfg *config.Config, repo *Repositories) (*Usecases, error) {
	authUsecase, err := auth.NewUsecase(cfg, repo.Auth)
	if err != nil {
		return nil, err
	}
	return &Usecases{auth: authUsecase}, nil
}
