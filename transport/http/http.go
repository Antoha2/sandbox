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
	GetUsers(user service.User) error
	DelUser(id int) error
	AddUser(user service.User) error
	UpdateUser(user service.User) error
}

type webImpl struct {
	service Service
	server  *http.Server
	log     *slog.Logger
	cfg     *config.Config
	port    int
}

func NewWeb(service Service, log *slog.Logger, cfg *config.Config) *webImpl {
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
