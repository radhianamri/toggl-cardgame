package response

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type response struct {
	Status string      `json:"status,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func JSONWithData(c echo.Context, statusCode int, msg string, data interface{}) error {
	return c.JSON(statusCode, response{
		Status: http.StatusText(statusCode),
		Msg:    msg,
		Data:   data,
	})
}

func BadRequest(c echo.Context, msg string) error {
	return JSONWithData(c, http.StatusBadRequest, msg, nil)
}

func InternalServerError(c echo.Context) error {
	return JSONWithData(c, http.StatusInternalServerError, "Internal Server Error", nil)
}

func NotFound(c echo.Context, msg string) error {
	return JSONWithData(c, http.StatusNotFound, msg, nil)
}

func ErrOrNotFound(c echo.Context, err error, msgNotFound string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, msgNotFound)
	}
	return c.String(http.StatusInternalServerError, "Internal Server Error")
}
