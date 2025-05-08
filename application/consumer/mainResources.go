package consumer

import (
	"com.github/confusionhill-aqw-ps/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Resources struct {
	Db *sqlx.DB
}

func NewResources(cfg *config.Config) (*Resources, error) {
	db, err := sqlx.Connect(cfg.Server.Database.Type, cfg.Server.Database.Host)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	return &Resources{
		Db: db,
	}, nil
}
