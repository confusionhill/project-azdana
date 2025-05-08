package application

import (
	"com.github/confusionhill-aqw-ps/application/consumer"
	webbackend "com.github/confusionhill-aqw-ps/application/webBackend"
	"com.github/confusionhill-aqw-ps/internal/config"
)

func RunApplication(cfg *config.Config) error {
	rsc, err := consumer.NewResources(cfg)
	if err != nil {
		return err
	}
	repo, err := consumer.NewRepositories(cfg, rsc)
	if err != nil {
		return err
	}
	usecases, err := consumer.NewUsecases(cfg, repo)
	if err != nil {
		return err
	}
	handlers, err := consumer.NewHandlers(cfg, usecases)
	if err != nil {
		return err
	}
	//Setup(cfg, rsc)
	return webbackend.RunWebBackendApp(cfg, handlers)
}
