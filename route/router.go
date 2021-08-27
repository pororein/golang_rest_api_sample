package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"repos.tdctechsupport.com/t2015014/user_management_server/api"
	"repos.tdctechsupport.com/t2015014/user_management_server/conf"
	"repos.tdctechsupport.com/t2015014/user_management_server/db"
	mw "repos.tdctechsupport.com/t2015014/user_management_server/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
	}))

	config := new(conf.AppConfig)
	config.GetConfig()
	e.Use(mw.MongoClientHandler(db.ConnectDB(config.MongoURL)))

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${level}")

	e.GET("/users/:email", api.GetUser)
	e.POST("/users/register", api.CreateUser)
	e.DELETE("/users/:email", api.DeleteUser)
	e.POST("/users/update/:email", api.UpdateUser)

	return e
}
