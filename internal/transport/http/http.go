package transport

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Antoha2/sandbox/internal/config"
	"github.com/Antoha2/sandbox/internal/service"
)

type Service interface {
	GetUsers(ctx context.Context, filter *service.QueryUsersFilter) ([]*service.User, error)
	GetUser(ctx context.Context, id int) (*service.User, error)
	DeleteUser(ctx context.Context, id int) (*service.User, error)          //user
	AddUser(ctx context.Context, user *service.User) (*service.User, error) //user   //sql returning
	UpdateUser(ctx context.Context, user *service.User) (*service.User, error)
}

type apiImpl struct {
	cfg     *config.Config
	log     *slog.Logger
	service Service
	server  *http.Server
}

// NewAPI
func NewApi(cfg *config.Config, log *slog.Logger, service Service) *apiImpl {
	return &apiImpl{
		service: service,
		log:     log,
		cfg:     cfg,
	}
}
