package auth

import (
	"context"
	"errors"

	"com.github/confusionhill-aqw-ps/internal/config"
	"com.github/confusionhill-aqw-ps/internal/model/dto/auth"
	"com.github/confusionhill-aqw-ps/internal/model/entity/game"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type Repository struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewRepository(cfg *config.Config, db *sqlx.DB) (*Repository, error) {
	return &Repository{cfg: cfg, db: db}, nil
}

func (r *Repository) createUser(ctx context.Context, req *auth.RegisterUserRequestDTO) (int64, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return 0, err
	}
	var user game.User
	tx.GetContext(ctx, &user, "SELECT id FROM users WHERE username = ? LIMIT 1", req.Username)
	if user.ID != 0 {
		tx.Rollback()
		return 0, errors.New("user already exists")
	}
	var newUser game.User
	err = tx.GetContext(ctx, &newUser, "INSERT INTO users (username, password, email, age, dob, signupip, gender, currentclass, plaColorEyes, plaColorHair, plaColorSkin, hairID, curServer) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id;", req.Username, req.Password, req.Email, req.Age, req.DOB, "", req.Gender, req.ClassID, req.EyeColor, req.HairColor, req.SkinColor, req.HairID, "WQW")
	if err != nil || newUser.ID == 0 {
		log.Error(err)
		tx.Rollback()
		return 0, errors.New("user already exists")
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO items (itemid, userid, equipped, sES, iLvl) VALUES (?, ?, ?, ?, ?)", 1, newUser.ID, 1, "Weapon", 1)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return 0, errors.New("user already exists")
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO items (itemid, userid, equipped, sES, iLvl, classXP, className) VALUES (?, ?, ?, ?, ?, ?, ?)", 16, newUser.ID, 1, "ar", 1, 0, "Warrior Class")
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return 0, errors.New("user already exists")
	}
	tx.Commit()
	return newUser.ID, nil
}

func (r *Repository) LoginUser(ctx context.Context, req auth.LoginUserRequestDTO) (*game.User, error) {
	var user game.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = ? AND password = ? LIMIT 1", req.Username, req.Password)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &user, nil
}
