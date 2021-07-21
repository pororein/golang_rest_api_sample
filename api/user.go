package api

import (
	model "repos.tdctechsupport.com/t2015014/user_management_server/model"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		email := c.Param("email")
		client := c.Get("MongoClient").(*mongo.Client)
		user := new(model.User)
		if err := user.Load(client, email); err != nil {
			return echo.NewHTTPError(
				fasthttp.StatusNotFound, "Member does not exists.")
		}
		return c.JSON(fasthttp.StatusOK, user)
	}
}
