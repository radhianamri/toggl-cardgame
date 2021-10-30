package database

import (
	"time"

	"github.com/radhianamri/toggl-cardgame/lib/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Dsn              string        `toml:"dsn"`
	MaxIdleTimeInMin time.Duration `toml:"max_idle_time_in_min"`
	MaxIdleConn      int           `toml:"max_idle_conn"`
	MaxOpenConn      int           `toml:"max_open_conn"`
}

func Init(cfg Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetConnMaxIdleTime(cfg.MaxIdleTimeInMin * time.Minute)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Info("MySQL succesfuly initialized")
	return db
}
