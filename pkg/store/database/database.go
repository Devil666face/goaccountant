package database

import (
	"log"

	"github.com/Devil666face/goaccountant/pkg/config"

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
		log.Fatalln(err)
	}
	return &d
}

func (d *Database) Migrate() error {
	return d.db.AutoMigrate(d.tables...)
}

func (d *Database) connect() error {
	if d.config.PostgresUse {
		return d.NewPsql()
	}
	return d.NewSqlite()
}
