package repository

import (
	"github.com/jmoiron/sqlx"
)

type Rep struct {
	DB *sqlx.DB
}

func NewRep(dbx *sqlx.DB) *Rep {
	return &Rep{
		DB: dbx,
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
