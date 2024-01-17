package config

import (
	"os"
)

type Config struct {
	//Env             string     `yaml:"env" env-default:"local"`
	HTTP     HTTPConfig `yaml:"http"`
	DBConfig DBConfig   `yaml:"postgres"`
	//MigrationsPath  string
	AddrAge         string `yaml:"addrage"`
	AddrGender      string `yaml:"addrgender"`
	AddrNationality string `yaml:"addrnationality"`
}

type DBConfig struct {
	User     string `yaml:"dbuser"`
	Password string `yaml:"dbpassword"`
	Host     string `yaml:"dbhost"`
	Port     string `yaml:"dbport"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"dbsslmode"`
}

type HTTPConfig struct {
	HostAddr string `yaml:"httpport"`
}

func NewCfg() *Config {
	return &Config{
		HTTP: HTTPConfig{
			HostAddr: os.Getenv("HTTP_PORT"),
		},
		DBConfig: DBConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Dbname:   os.Getenv("DB_NAME"),
			Sslmode:  os.Getenv("DB_SSLMODE"),
		},
		AddrAge:         os.Getenv("ADDR_AGE"),
		AddrGender:      os.Getenv("ADDR_GENDER"),
		AddrNationality: os.Getenv("ADDR_NATIONALITY"),
	}
}
