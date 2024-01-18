package transport

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/Antoha2/sandbox/config"
	"github.com/Antoha2/sandbox/service"
)

type Service interface {
	GetUsers(ctx context.Context, filter *service.GetQueryFilter) ([]*service.User, error)
	GetUser(ctx context.Context, id int) (*service.User, error)
	DelUser(ctx context.Context, id int) error
	AddUser(ctx context.Context, user *service.User) (*service.User, error)
	UpdateUser(ctx context.Context, user *service.User) (*service.User, error)
}

type webImpl struct {
	cfg     *config.Config
	log     *slog.Logger
	service Service
	server  *http.Server
	port    int
}

func NewWeb(cfg *config.Config, log *slog.Logger, service Service) *webImpl {
	HTTPport, err := strconv.Atoi(cfg.HTTP.HostAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &webImpl{
		service: service,
		log:     log,
		cfg:     cfg,
		port:    HTTPport,
	}
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
