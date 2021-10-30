package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Status string      `json:"status,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func ResponseJSONWithData(c echo.Context, statusCode int, msg string, data interface{}) error {
	return c.JSON(statusCode, response{
		Status: http.StatusText(statusCode),
		Msg:    msg,
		Data:   data,
	})
}

func ResponseOKWithData(c echo.Context, data interface{}) error {
	return ResponseJSONWithData(c, http.StatusOK, "OK", data)
}

func ResponseOK(c echo.Context) error {
	return ResponseJSONWithData(c, http.StatusOK, "OK", nil)
}

func ResponseUnauthorized(c echo.Context, msg string) error {
	return ResponseJSONWithData(c, http.StatusUnauthorized, msg, nil)
}

func ResponseBadRequest(c echo.Context, msg string) error {
	return ResponseJSONWithData(c, http.StatusBadRequest, msg, nil)
}

func ResponseInternalServerError(c echo.Context) error {
	return ResponseJSONWithData(c, http.StatusInternalServerError, "Internal Server Error", nil)
}

func ResponseNotFound(c echo.Context, msg string) error {
	return ResponseJSONWithData(c, http.StatusNotFound, msg, nil)
}
