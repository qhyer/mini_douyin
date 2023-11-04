package orm

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Config mysql config.
type Config struct {
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max lifetime.
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	if err != nil {
		log.Error("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(c.Idle)
	sqlDB.SetMaxOpenConns(c.Active)
	sqlDB.SetConnMaxLifetime(c.IdleTimeout / time.Second)
	return
}
