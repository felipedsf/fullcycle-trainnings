package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"sync"
)

type Migration struct {
	Db *sql.DB
}

var (
	instance *Migration
	once     sync.Once
)

func NewMigration(db *sql.DB) *Migration {
	once.Do(func() {
		instance = &Migration{Db: db}
	})
	return instance
}

func (m *Migration) AutoMigrate() error {
	logrus.Info("Auto migration started")
	logrus.Info("Creating table 'orders' if it not exists")
	_, err := m.Db.Exec("CREATE TABLE IF NOT EXISTS orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	if err != nil {
		logrus.Error("Error creating table 'orders': ", err)
		return err
	}
	logrus.Info("Auto migration finished")
	return nil
}
