package dbrepo

import (
	"database/sql"

	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/repository"
)

// holds two things: app and database connection pool
type MysqlDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewMysqlRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {

	return &MysqlDBRepo{
		App: a,
		DB:  conn,
	}
}
