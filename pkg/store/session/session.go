package session

import (
	"time"

	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

const (
	table = "storage"
)

type Store struct {
	Store         *session.Store
	Storage       fiber.Storage
	config        *config.Config
	sqliteConnect string
	psqlConnect   string
}

func New(cfg *config.Config, db *database.Database) *Store {
	s := Store{
		config:        cfg,
		sqliteConnect: db.SqliteConnect,
		psqlConnect:   db.PsqlConnect,
	}
	s.setStorage()
	s.Store = s.getStore()
	return &s
}

func (s *Store) setStorage() {
	if s.config.PostgresUse {
		s.Storage = s.NewPsqlStorage()
	}
	s.Storage = s.NewSqliteStorage()
}

func (s *Store) getStore() *session.Store {
	return session.New(session.Config{
		// CookieSecure: true,
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 5,
		Storage:        s.Storage,
	})
}
