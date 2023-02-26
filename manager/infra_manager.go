package manager

import (
	"kel1-stockbite-projects/config"

	"github.com/jmoiron/sqlx"
)

type InfraManager interface {
	SqlDb() *sqlx.DB
}

type infraManager struct {
	db     *sqlx.DB
	config config.Config
}

func (i *infraManager) SqlDb() *sqlx.DB {
	return i.db
}

func (i *infraManager) InitDb() {
	connectDb, err := sqlx.Open("postgres", i.config.DataSourceName)

	if err != nil {
		panic(err.Error())
	}
	CheckErrorConfig := connectDb.Ping()
	if CheckErrorConfig != nil {
		panic(CheckErrorConfig)
	}
	i.db = connectDb
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		config: config,
	}
	infra.InitDb()
	return &infra
}
