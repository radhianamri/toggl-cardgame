package config

import (
	"io/ioutil"
	"os"
	"time"

	toml "github.com/pelletier/go-toml"
	"github.com/radhianamri/toggl-cardgame/lib/database"
	"github.com/radhianamri/toggl-cardgame/lib/log"
)

type Config struct {
	Main       Main            `toml:"Main"`
	Middleware Middleware      `toml:"Middleware"`
	DB         database.Config `toml:"Database"`
	Rest       Rest            `toml:"Rest"`
}

type Main struct {
	Env      string `toml:"env"`
	Timezone string `toml:"timezone"`
}

type Middleware struct {
	LogFormat     string `toml:"log_format"`
	LogTimeFormat string `toml:"log_time_format"`
	GzipLevel     int    `toml:"gzip_level"`
}

type Rest struct {
	Port         string        `toml:"port"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

func Init() *Config {
	f, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var config Config
	if err = toml.Unmarshal(f, &config); err != nil {
		log.Fatalf("failed to unmarhsal config: %v", err)
	}
	timezoneInit(config.Main.Timezone)
	log.Info("Config succesfuly initialized")
	return &config
}

func timezoneInit(tz string) {
	if tz == "" {
		tz = "Asia/Jakarta"
	}
	os.Setenv("TZ", tz)
}
