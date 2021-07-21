package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"repos.tdctechsupport.com/t2015014/user_management_server/api"
	"repos.tdctechsupport.com/t2015014/user_management_server/conf"
	"repos.tdctechsupport.com/t2015014/user_management_server/db"
	"repos.tdctechsupport.com/t2015014/user_management_server/handler"
	mw "repos.tdctechsupport.com/t2015014/user_management_server/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	config := new(conf.AppConfig)

	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler
	e.Use(mw.MongoClientHandler(db.ConnectDB(config.MongoURL)))

	e.GET("/users/:email", api.GetUser())

	return e
}
