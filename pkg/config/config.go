package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	https = "https"
)

type Config struct {
	Ip               string `env:"IP" env-default:"127.0.0.1"`
	PortHttp         uint   `env:"PORT_HTTP" env-default:"8000"`
	PortHttps        uint   `env:"PORT_HTTPS" env-default:"4443"`
	SqliteDB         string `env:"SQLITE_DB" env-default:"db.sqlite3"`
	AllowHost        string `env:"ALLOW_HOST" env-default:"localhost"`
	UseTls           bool   `env:"TLS" env-default:"false"`
	TlsKey           string `env:"TLS_KEY" env-default:"server.key"`
	TlsCrt           string `env:"TLS_CERT" env-default:"server.crt"`
	PostgresUse      bool   `env:"POSTGRES" env-default:"false"`
	PostgresHost     string `env:"POSTGRES_HOST" env-default:"localhost"`
	PostgresPort     string `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresDb       string `env:"POSTGRES_DB" env-default:"db"`
	PostgresUser     string `env:"POSTGRES_USER" env-default:"superuser"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" env-default:"Qwerty123!"`
	ConnectHttp      string
	ConnectHttps     string
	HttpsRedirect    string
}

func New() *Config {
	cfg := Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}
	cfg.ConnectHttp = fmt.Sprintf("%v:%v", cfg.Ip, cfg.PortHttp)
	cfg.ConnectHttps = fmt.Sprintf("%v:%v", cfg.Ip, cfg.PortHttps)
	cfg.HttpsRedirect = fmt.Sprintf("%s://%s:%d", https, cfg.AllowHost, cfg.PortHttps)
	return &cfg
}
