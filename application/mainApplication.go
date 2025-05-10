package application

import (
	"com.github/confusionhill-aqw-ps/application/consumer"
	"com.github/confusionhill-aqw-ps/internal/config"
)

func RunApplication(cfg *config.Config) (*consumer.Resources, *consumer.Repositories, *consumer.Usecases, *consumer.Handlers, error) {
	rsc, err := consumer.NewResources(cfg)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	repo, err := consumer.NewRepositories(cfg, rsc)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	usecases, err := consumer.NewUsecases(cfg, repo)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	handlers, err := consumer.NewHandlers(cfg, usecases)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return rsc, repo, usecases, handlers, nil
}
