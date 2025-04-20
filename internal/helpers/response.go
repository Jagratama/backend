package helpers

import "github.com/labstack/echo/v4"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponseHTTP(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, Response{
		Message: message,
		Data:    data,
	})
}
