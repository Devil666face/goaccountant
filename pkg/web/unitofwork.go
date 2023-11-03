package web

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"

	"github.com/gofiber/fiber/v2"
	fibersession "github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type Uof struct {
	viewctx    *ViewCtx
	database   *database.Database
	config     *config.Config
	session    *session.Store
	ctxsession *fibersession.Session
}

func NewUof(ctx *fiber.Ctx, db *database.Database, cfg *config.Config, s *session.Store) *Uof {
	return &Uof{
		viewctx:  NewViewCtx(ctx),
		database: db,
		config:   cfg,
		session:  s,
	}
}

func (uof *Uof) ViewCtx() *ViewCtx {
	return uof.viewctx
}

func (uof *Uof) FiberCtx() *fiber.Ctx {
	return uof.viewctx.Ctx
}

func (uof *Uof) Database() *gorm.DB {
	return uof.database.DB()
}

func (uof *Uof) Store() *fibersession.Store {
	return uof.session.Store()
}

func (uof *Uof) Storage() fiber.Storage {
	return uof.session.Storage()
}

func (uof *Uof) Config() *config.Config {
	return uof.config
}

func (uof *Uof) GetSession() error {
	var err error
	if uof.ctxsession, err = uof.Store().Get(uof.FiberCtx()); err != nil {
		return err
	}
	return nil
}

func (uof *Uof) SetInSession(key string, val any) {
	uof.ctxsession.Set(key, val)
}

func (uof *Uof) SaveSession() error {
	return uof.ctxsession.Save()
}
