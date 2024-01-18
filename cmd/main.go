package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Antoha2/sandbox/cmd/logger"
	"github.com/Antoha2/sandbox/config"
	providerAge "github.com/Antoha2/sandbox/providers/age"
	providerGender "github.com/Antoha2/sandbox/providers/gender"
	providerNat "github.com/Antoha2/sandbox/providers/nationality"
	"github.com/Antoha2/sandbox/repository"
	"github.com/Antoha2/sandbox/service"
	transport "github.com/Antoha2/sandbox/transport/http"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func init() {
	// убрать это из инит и перенесть мастлоад
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	Run()
}

func Run() {

	cfg := config.MustLoad()
	slog := logger.SetupLogger(cfg.Env)
	dbx, err := initDb(cfg)
	if err != nil { // для единообразия лучше проверку убрать в функцию, и назвать ее мастинитдб
		fmt.Println(err)
		os.Exit(1)
	}

	rep := repository.NewRep(slog, dbx)
	// переменные здесь должны быть существительными
	getAge := providerAge.NewGetAge()
	getGender := providerGender.NewGetGender()
	getNat := providerNat.NewGetNat()
	serv := service.NewServ(cfg, slog, rep, getAge, getGender, getNat) //, getAge
	trans := transport.NewWeb(cfg, slog, serv)

	go trans.StartHTTP()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	trans.Stop()

}

func initDb(cfg *config.Config) (*sqlx.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Dbname,
		cfg.DBConfig.Sslmode,
	)

	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		// добавлять инфу в ошибку лучше через errors.Wrap()
		return nil, fmt.Errorf("1 failed to parse config: %v", err) // что значит 1)??
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf("2 failed to create connection db: %v", err) // что значит 2)?? плохая ошибка, нужна конкретика
	}

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf("4 error to ping connection pool: %v", err) // через wrap, здесь и во всех других местах
	}
	// все принты должны быть через логер slog
	log.Printf("Подключение к базе данных на http://127.0.0.1:%v\n", cfg.DBConfig.Port)
	return dbx, nil
}
