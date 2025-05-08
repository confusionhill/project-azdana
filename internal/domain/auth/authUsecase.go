package auth

import (
	"context"

	"com.github/confusionhill-aqw-ps/internal/config"
	"com.github/confusionhill-aqw-ps/internal/model/dto/auth"
	"com.github/confusionhill-aqw-ps/internal/model/entity/game"
)

type Usecase struct {
	cfg  *config.Config
	repo *Repository
}

func NewUsecase(cfg *config.Config, repo *Repository) (*Usecase, error) {
	return &Usecase{cfg: cfg, repo: repo}, nil
}

func (u *Usecase) registerUserUsecase(ctx context.Context, req *auth.RegisterUserRequestDTO) (int64, error) {
	return u.repo.createUser(ctx, req)
}

func (u *Usecase) loginUserUsecase(ctx context.Context, req auth.LoginUserRequestDTO) (*game.User, error) {
	user, err := u.repo.loginUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
