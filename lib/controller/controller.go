package controller

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/radhianamri/toggl-cardgame/lib/log"
	"github.com/radhianamri/toggl-cardgame/lib/validator"
)

func ReadRequestBody(c echo.Context, model interface{}) error {
	if err := c.Bind(&model); err != nil {
		log.Error(err)
		return errors.New("invalid request body")
	}
	if err := validator.ValidateStructInput(model); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
