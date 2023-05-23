package database

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func Initialize(dsn, driver string, maxOpenConns, maxIdleConns int) {
	once.Do(func() {
		var err error

		db, err = sqlx.Open(driver, dsn)

		if err != nil {
			panic(err)
		}

		db.SetMaxIdleConns(maxIdleConns)
		db.SetMaxOpenConns(maxOpenConns)
	})
}

// GetDB 获得SQLX的DB实例
func GetDB() *sqlx.DB {

	if db == nil {
		panic("Database Not Initialize.")
	}

	err := db.Ping()
	if err != nil {
		panic("Database Connection Failed.")
	}

	return db
}
