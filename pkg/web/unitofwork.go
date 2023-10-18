package web

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"

	"github.com/gofiber/fiber/v2"
	fiberses "github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type Uof struct {
	ctx *Ctx
	db  *database.Database
	cfg *config.Config
	ses *session.Store
}

func NewUof(ctx *fiber.Ctx, db *database.Database, cfg *config.Config, ses *session.Store) *Uof {
	return &Uof{
		ctx: NewCtx(ctx),
		db:  db,
		cfg: cfg,
		ses: ses,
	}
}

func (uof *Uof) Ctx() *Ctx {
	return uof.ctx
}

func (uof *Uof) FiberCtx() *fiber.Ctx {
	return uof.ctx.Ctx
}

func (uof *Uof) Database() *gorm.DB {
	return uof.db.DB()
}

func (uof *Uof) Store() *fiberses.Store {
	return uof.ses.Store()
}

func (uof *Uof) Storage() fiber.Storage {
	return uof.ses.Storage()
}

func (uof *Uof) Config() *config.Config {
	return uof.cfg
}
