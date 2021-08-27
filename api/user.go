package api

import (
	mw "repos.tdctechsupport.com/t2015014/user_management_server/middleware"
	model "repos.tdctechsupport.com/t2015014/user_management_server/model"

	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(c echo.Context) error {
	email := c.Param("email")
	c.Logger().Debug("email:", email)
	client := c.Get(mw.ClientKey).(*mongo.Client)
	user := new(model.User)
	if err := user.Load(client, email); err != nil {
		return echo.NewHTTPError(
			http.StatusNotFound, "Member does not exists.")
	}
	c.Logger().Debug("user:", user)
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := model.User{
		ID:        primitive.NewObjectID(),
		EMail:     "",
		FirstName: "",
		LastName:  "",
	}
	if err := c.Bind(&user); err != nil {
		c.Logger().Error("request parameter analyze error:", err)
		return echo.NewHTTPError(
			http.StatusBadRequest, "Invalid Parameter.")
	}
	client := c.Get(mw.ClientKey).(*mongo.Client)
	result, err := user.Save(client)
	if err != nil {
		c.Logger().Error("DB Insert error:", err)
		return echo.NewHTTPError(
			http.StatusInternalServerError, "DB Insert Failed.")
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	email := c.Param("email")
	c.Logger().Debug("email:", email)
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.Logger().Error("request parameter analyze error:", err)
		return echo.NewHTTPError(
			http.StatusBadRequest, "Invalid Parameter.")
	}
	client := c.Get(mw.ClientKey).(*mongo.Client)
	result, err := user.Update(client, email)
	if err != nil {
		c.Logger().Error("DB Update error:", err)
		return echo.NewHTTPError(
			http.StatusInternalServerError, "DB Update Failed.")
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	email := c.Param("email")
	c.Logger().Debug("email:", email)
	client := c.Get(mw.ClientKey).(*mongo.Client)
	user := new(model.User)
	result, err := user.Delete(client, email)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusNotFound, "DB Data Delete Failed.")
	}
	c.Logger().Debug("user:", result)
	return c.JSON(http.StatusOK, result)
}
