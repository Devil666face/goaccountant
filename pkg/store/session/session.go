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
	store         *session.Store
	storage       fiber.Storage
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
	if s.config.PostgresUse {
		s.storage = s.NewPsqlStorage()
	}
	s.storage = s.NewSqliteStorage()
	s.store = s.getStore()
	return &s
}

func (s *Store) Store() *session.Store {
	return s.store
}

func (s *Store) Storage() fiber.Storage {
	return s.storage
}

func (s *Store) getStore() *session.Store {
	return session.New(session.Config{
		// CookieSecure: true,
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 5,
		Storage:        s.storage,
	})
}
