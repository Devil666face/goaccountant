package session

import (
	"github.com/gofiber/storage/postgres/v3"
)

func (s *SessionStore) NewPsqlStorage() *postgres.Storage {
	return postgres.New(postgres.Config{
		ConnectionURI: s.psqlConnect,
		Table:         table,
		Reset:         false,
	})
}
