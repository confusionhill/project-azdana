package consumer

import (
	"com.github/confusionhill-aqw-ps/internal/config"
	"com.github/confusionhill-aqw-ps/internal/domain/auth"
)

type Repositories struct {
	auth *auth.Repository
}

func NewRepositories(cfg *config.Config, rsc *Resources) (*Repositories, error) {
	authRepo, err := auth.NewRepository(cfg, rsc.Db)
	if err != nil {
		return nil, err
	}
	return &Repositories{
		auth: authRepo,
	}, nil
}
