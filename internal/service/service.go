package service

import (
	"context"
	"log/slog"

	"github.com/Antoha2/catAuto/internal/config"
	"github.com/Antoha2/catAuto/internal/repository"
)

const DefaultPropertyAge = 0
const DefaultPropertyOffset = 0
const DefaultPropertyLimit = 100

type Repository interface {
	GetUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context, filter *repository.RepQueryFilter) ([]*repository.RepUser, error)
	AddUser(ctx context.Context, user *repository.RepUser) (int, error)
	DeleteUser(ctx context.Context, id int) (*repository.RepUser, error)
	UpdateUser(ctx context.Context, user *repository.RepUser) (*repository.RepUser, error)
}

type AgeProvider interface {
	GetAge(ctx context.Context, name string) (int, error)
}
type GenderProvider interface {
	GetGender(ctx context.Context, name string) (string, error)
}
type NationalityProvider interface {
	GetNationality(ctx context.Context, name string) (string, error)
}

type servImpl struct {
	cfg               *config.Config
	log               *slog.Logger
	rep               *repository.Rep
	ageClient         AgeProvider
	genderClient      GenderProvider
	nationalityClient NationalityProvider
}

func NewServ(
	cfg *config.Config,
	log *slog.Logger,
	rep *repository.Rep,
	ageClient AgeProvider,
	genderClient GenderProvider,
	nationalityClient NationalityProvider) *servImpl {
	return &servImpl{
		rep:               rep,
		log:               log,
		cfg:               cfg,
		ageClient:         ageClient,
		genderClient:      genderClient,
		nationalityClient: nationalityClient,
	}
}

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	SurName     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type QueryUsersFilter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	SurName     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}
