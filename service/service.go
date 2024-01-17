package service

import (
	"log/slog"

	"github.com/Antoha2/sandbox/config"
	"github.com/Antoha2/sandbox/repository"
)

type Repository interface {
	AddUser(user repository.RepUser) error
	//UserSaver(ctx context.Context, email string, passHash []byte) (uid int64, err error)
	//UserProvider(ctx context.Context, email string) (models.User, error)
}

type Query struct {
	Name string
	Addr string
}

type AgeProvider interface {
	GetAge(request *Query) (int, error)
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

func NewServ(Rep *repository.Rep,
	cfg *config.Config,
	log *slog.Logger,
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

type GetQueryUser struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	SurName     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Page        int
}
