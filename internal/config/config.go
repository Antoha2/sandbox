package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// убрать теги для ямла, его нет в проекте
type Config struct {
	Env      string     `yaml:"env" env:"ENV" env-default:"local"`
	HTTP     HTTPConfig `yaml:"http" `
	DBConfig DBConfig   `yaml:"postgres"`
	//MigrationsPath  string удалить
	// во всех местах проекта, Addr переименоват в URL
	AddrAge         string `yaml:"addrage" env:"ADDR_AGE"`                 // URLProviderAge/URL_PROVIDER_AGE
	AddrGender      string `yaml:"addrgender" env:"ADDR_GENDER"`           // по примемеру выше
	AddrNationality string `yaml:"addrnationality" env:"ADDR_NATIONALITY"` // по примеру
}

type DBConfig struct {
	User     string `yaml:"dbuser" env:"DB_USER" env-default:"user"`
	Password string `yaml:"dbpassword" env:"DB_PASSWORD" env-default:"user"`
	Host     string `yaml:"dbhost" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"dbport" env:"DB_PORT" env-default:"5432"`
	Dbname   string `yaml:"dbname" env:"DB_DBNAME" env-default:"test"`
	Sslmode  string `yaml:"dbsslmode" env:"DB_SSLMODE" env-default:""`
}

type HTTPConfig struct {
	HostAddr string `yaml:"httpport" env:"HTTP_PORT" env-default:"8080"` // HostAddr переименовать в Port
}

// загрузка конфига из .env
func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
