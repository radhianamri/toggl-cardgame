package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radhianamri/toggl-cardgame/lib/config"
	"github.com/radhianamri/toggl-cardgame/lib/database"
	"github.com/radhianamri/toggl-cardgame/lib/log"
	"github.com/radhianamri/toggl-cardgame/lib/validator"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func init() {
	log.Init()
	validator.Init()
}

func main() {
	cfg := config.Init()
	_ = database.Init(cfg.DB)

	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           cfg.Middleware.LogFormat,
			CustomTimeFormat: cfg.Middleware.LogTimeFormat,
		}),
		middleware.GzipWithConfig(middleware.GzipConfig{Level: cfg.Middleware.GzipLevel}),
	)

	server := &http.Server{
		Addr:         cfg.Rest.Port,
		ReadTimeout:  cfg.Rest.ReadTimeout,
		WriteTimeout: cfg.Rest.WriteTimeout,
		Handler:      e,
	}

	log.Info("Server running...")
	server.ListenAndServe()
}
