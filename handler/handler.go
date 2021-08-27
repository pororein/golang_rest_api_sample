package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := fasthttp.StatusInternalServerError
	msg := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	if !c.Response().Committed {
		c.JSON(code, map[string]interface{}{
			"statusCode": code,
			"message":    msg,
		})
	}
}
