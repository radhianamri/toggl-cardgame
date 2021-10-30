package db

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

var dbConn *gorm.DB

func Init(cfg Config) {
	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	db.Set("gorm:auto_preload", true)
	db.Callback().Update().Remove("gorm:update_time_stamp")

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
	dbConn = db
}

func GetConn() *gorm.DB {
	return dbConn
}

func SetConn() *gorm.DB {
	return dbConn
}
