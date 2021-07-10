package api

import (
	"main/model"
)

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		email, _ := c.Param("email")
		client := c.Get("MongoClient").(*mongo.Client)
		user := new(model.User)
		if err := user.Load(client, email); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "Member does not exists.")
		}
		return c.JSON(fasthttp.StatusOK, user)
	}
}
