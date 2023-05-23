package database

import (
	"sync"

	"github.com/jmoiron/sqlx"

	// 单元测试sqlx使用的是sqlite驱动
	_ "github.com/mattn/go-sqlite3"

	// 程序要使用mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sqlx.DB
	config DBConfig
	once   sync.Once
)

type DBConfig struct {
	// 这些是db的配置
	dsn          string
	driver       string
	maxOpenConns int
	maxIdleConns int
}

func Initialize(dsn, driver string, maxOpenConns, maxIdleConns int) {
	config = DBConfig{
		dsn:          dsn,
		driver:       driver,
		maxOpenConns: maxOpenConns,
		maxIdleConns: maxIdleConns,
	}
}

func connect() {
	once.Do(func() {
		var err error

		if db, err = sqlx.Open(config.driver, config.dsn); err != nil {
			panic(err)
		}

		db.SetMaxIdleConns(config.maxIdleConns)
		db.SetMaxOpenConns(config.maxOpenConns)
	})
}

// GetDB 获得SQLX的DB实例
func GetDB() *sqlx.DB {

	if db == nil {
		connect()
	}

	err := db.Ping()
	if err != nil {
		panic("Database Connection Failed.")
	}

	return db
}
