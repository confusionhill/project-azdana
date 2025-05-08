package application

import "github.com/jmoiron/sqlx"

type Resources struct {
	Db *sqlx.DB
}

func NewResources() (*Resources, error) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/aqw")
	return &Resources{
		Db: db,
	}, err
}
