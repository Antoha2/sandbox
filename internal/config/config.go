package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env            string `env:"ENV" env-default:"local"`
	HTTP           HTTPConfig
	DBConfig       DBConfig
	URLAge         string `env:"URL_AGE"`
	URLGender      string `env:"URL_GENDER"`
	URLNationality string `env:"URL_NATIONALITY"`
}

type DBConfig struct {
	User     string `env:"DB_USER" env-default:"user"`
	Password string `env:"DB_PASSWORD" env-default:"user"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	Dbname   string `env:"DB_DBNAME" env-default:"test"`
	Sslmode  string `env:"DB_SSLMODE" env-default:""`
}

type HTTPConfig struct {
	HostPort string `env:"HTTP_PORT" env-default:"8080"`
}

//загрузка конфига из .env
func MustLoad() *Config {

	if err := godotenv.Load(); err != nil {
		panic("No .env file found" + err.Error())
	}
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
