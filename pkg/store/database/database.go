package database

import (
	"log"

	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/utils"
	"github.com/Devil666face/goaccountant/pkg/web/models"

	"gorm.io/gorm"
)

// var DB *gorm.DB

type Database struct {
	db            *gorm.DB
	config        *config.Config
	tables        []any
	SqliteConnect string
	PsqlConnect   string
}

func New(cfg *config.Config, tables []any) *Database {
	d := Database{
		config: cfg,
		tables: tables,
	}
	if err := d.connect(); err != nil {
		//nolint:revive //If database not open - close app
		log.Fatalln(err)
	}
	if err := d.migrate(); err != nil {
		log.Print(err)
	}
	if err := d.createSuperuser(); err != nil {
		log.Print(err)
	}
	return &d
}

func (d *Database) DB() *gorm.DB {
	return d.db
}

func (d *Database) migrate() error {
	return d.db.AutoMigrate(d.tables...)
}

func (d *Database) connect() error {
	if d.config.PostgresUse {
		return d.NewPsql()
	}
	return d.NewSqlite()
}

func (d *Database) createSuperuser() error {
	hash, err := utils.GenHash(d.config.SuperuserPassword)
	if err != nil {
		return err
	}
	u := models.User{
		Username: d.config.Superuser,
		Admin:    true,
		Password: hash,
	}
	if err := u.Create(d.db); err != nil {
		return err
	}
	return nil
}
