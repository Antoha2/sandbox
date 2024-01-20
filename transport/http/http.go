package transport

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Antoha2/sandbox/internal/config"
	"github.com/Antoha2/sandbox/internal/service"
)

// константы в отдельный фаил
const ID = "id"
const AGE = "age"
const LIMIT = "limit"
const OFFSET = "offset"

type Service interface {
	GetUsers(ctx context.Context, filter *service.QueryUsersFilter) ([]*service.User, error)
	GetUser(ctx context.Context, id int) (*service.User, error)
	DelUser(ctx context.Context, id int) (*service.User, error)             //user
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
func NewWeb(cfg *config.Config, log *slog.Logger, service Service) *apiImpl {

	//HTTPport, err := strconv.Atoi(cfg.HTTP.HostAddr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	return &apiImpl{
		service: service,
		log:     log,
		cfg:     cfg,
	}
}

// перенеси ближе к функции старта
func (a *apiImpl) Stop() {
	if err := a.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
