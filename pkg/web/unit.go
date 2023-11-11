package web

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/Devil666face/goaccountant/pkg/web/validators"

	"github.com/gofiber/fiber/v2"
	fibersession "github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type Unit struct {
	viewctx    *ViewCtx
	database   *database.Database
	config     *config.Config
	session    *session.Store
	validator  *validators.Validator
	ctxsession *fibersession.Session
}

func NewUnit(
	c *fiber.Ctx,
	_database *database.Database,
	_config *config.Config,
	_session *session.Store,
	_validator *validators.Validator,
) *Unit {
	return &Unit{
		viewctx:   NewViewCtx(c),
		database:  _database,
		config:    _config,
		session:   _session,
		validator: _validator,
	}
}

func (unit *Unit) ViewCtx() *ViewCtx {
	return unit.viewctx
}

func (unit *Unit) Ctx() *fiber.Ctx {
	return unit.viewctx.Ctx
}

func (unit *Unit) Database() *gorm.DB {
	return unit.database.DB()
}

func (unit *Unit) Store() *fibersession.Store {
	return unit.session.Store()
}

func (unit *Unit) Storage() fiber.Storage {
	return unit.session.Storage()
}

func (unit *Unit) Config() *config.Config {
	return unit.config
}

func (unit *Unit) Validator() *validators.Validator {
	return unit.validator
}

func (unit *Unit) getSession() error {
	var err error
	if unit.ctxsession, err = unit.Store().Get(unit.Ctx()); err != nil {
		return err
	}
	return nil
}

func (unit *Unit) SetInSession(key string, val any) error {
	if err := unit.getSession(); err != nil {
		return err
	}
	unit.ctxsession.Set(key, val)
	return unit.SaveSession()
}

func (unit *Unit) GetFromSession(key string) (any, error) {
	if err := unit.getSession(); err != nil {
		return nil, err
	}
	return unit.ctxsession.Get(key), nil
}

func (unit *Unit) SaveSession() error {
	return unit.ctxsession.Save()
}

func (unit *Unit) DestroySession() error {
	return unit.ctxsession.Destroy()
}

func (unit *Unit) DestroySessionByID(sessID string) error {
	if sessID == "" {
		return nil
	}
	return unit.Store().Delete(sessID)
}

func (unit *Unit) SessionID() (string, error) {
	if err := unit.getSession(); err != nil {
		return "", err
	}
	return unit.ctxsession.ID(), nil
}
