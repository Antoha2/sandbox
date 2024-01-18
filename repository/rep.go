package repository

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type Rep struct {
	log *slog.Logger
	DB  *sqlx.DB
}

func NewRep(log *slog.Logger, dbx *sqlx.DB) *Rep {
	return &Rep{
		log: log,
		DB:  dbx,
	}
}

type RepUser struct {
	Id          int
	Name        string
	SurName     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

type RepQueryFilter struct {
	Id          int
	Name        string
	SurName     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
	Offset      int
	Limit       int
}
