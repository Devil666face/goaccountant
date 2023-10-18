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
	ctx      *Ctx
	database *database.Database
	config   *config.Config
	session  *session.Store
}

func NewUof(c *fiber.Ctx, db *database.Database, cfg *config.Config, s *session.Store) *Uof {
	return &Uof{
		ctx:      NewCtx(c),
		database: db,
		config:   cfg,
		session:  s,
	}
}

func (uof *Uof) Ctx() *Ctx {
	return uof.ctx
}

func (uof *Uof) FiberCtx() *fiber.Ctx {
	return uof.ctx.Ctx
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
