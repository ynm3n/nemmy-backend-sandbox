package internal

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	e.GET("/api/users/:userID", getUser)
}

func getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.String(http.StatusBadRequest, "userIDは整数にしてください🫠")
		return nil
	} else if id > 100 || id < 0 {
		c.String(http.StatusBadRequest, "0~99までのuserIDを指定してください🫠")
		return nil
	}

	// 本来ならここでDBにアクセスしてユーザー情報を取得する
	res := struct {
		UserID   int    `json:"userId"`
		UserName string `json:"username"`
	}{
		UserID:   id,
		UserName: usernames[id%4],
	}

	if err := c.JSON(http.StatusOK, res); err != nil {
		return err
	}
	return nil
}

var usernames = [4]string{
	"nemmy",
	"light",
	"ynm",
	"bandou eiji",
}
