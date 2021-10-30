package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radhianamri/toggl-cardgame/lib/config"
	db "github.com/radhianamri/toggl-cardgame/lib/database"
	"github.com/radhianamri/toggl-cardgame/lib/log"
	"github.com/radhianamri/toggl-cardgame/lib/validator"
	"github.com/radhianamri/toggl-cardgame/services/decks"
)

func init() {
	log.Init()
	validator.Init()
}

func main() {
	cfg := config.Init()
	db.Init(cfg.DB)

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
	v1 := e.Group("/v1")
	decks.RegisterRoutes(v1)

	log.Info("Server running...")
	server.ListenAndServe()
}
