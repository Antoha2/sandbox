package service

import (
	"context"
	"log/slog"

	"github.com/Antoha2/sandbox/config"
	"github.com/Antoha2/sandbox/repository"
)

// вообще все кроме cmd надо в папку интернел перенести

type Repository interface {
	GetUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context, filter *repository.RepQueryFilter) ([]*repository.RepUser, error)
	AddUser(ctx context.Context, user *repository.RepUser) (int, error)
	DelUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user *repository.RepUser) (*repository.RepUser, error)
	//UserSaver(ctx context.Context, email string, passHash []byte) (uid int64, err error)
	//UserProvider(ctx context.Context, email string) (models.User, error)
}

type Query struct {
	Name string
	Addr string // это не должно быть здесь
}

type AgeProvider interface {
	GetAge(request *Query) (int, error) // первый аргумент всегда контекст
}
type GenderProvider interface {
	GetGender(request *Query) (string, error)
}
type NationalityProvider interface {
	GetNationality(request *Query) (string, error)
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
	Rep *repository.Rep, // c маленькой
	ageClient AgeProvider,
	genderClient GenderProvider,
	nationalityClient NationalityProvider) *servImpl {
	return &servImpl{
		rep:               Rep,
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

type GetQueryFilter struct { // QueryUsersFilter
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
