package cmd

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/urfave/cli"
	"github.com/sunmi-OS/gocore/api"
	"fmt"
	"monomer-service-demo/model"
	"github.com/jinzhu/gorm"
	"errors"
)

var ModelUser model.User

func RunApi(c *cli.Context) error {

	fmt.Println("API-Service开始运行，通过*******进行访问！")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.Any("/login", func(c echo.Context) error {

		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		request := api.NewRequest(c)
		response := api.NewResponse(c)
		username := request.GetParam("username").Require(true).GetString()
		password := request.GetParam("password").Require(true).GetString()

		err := request.GetError()
		if err != nil {
			return response.RetError(err, -1)
		}

		user, err := ModelUser.CheckUsername(username)
		if err != nil {
			return response.RetError(err, -1)
		}
		if user.Password != password {
			return response.RetError(errors.New("用户密码不正确"), -1)
		}

		response.SetMsg("登录成功")
		return response.RetSuccess(user.Id)
	})

	e.Any("/registered", func(c echo.Context) error {

		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		request := api.NewRequest(c)
		response := api.NewResponse(c)

		username := request.GetParam("username").Require(true).GetString()
		password := request.GetParam("password").Require(true).GetString()

		err := request.GetError()
		if err != nil {
			return response.RetError(err, -1)
		}

		// 查询用户是否存在
		user, err := ModelUser.CheckUsername(username)
		if err != nil && err != gorm.ErrRecordNotFound {
			return response.RetError(err, 400)
		}
		if user.Username != "" {
			return response.RetError(errors.New("用户名已存在"), -1)
		}

		// 创建用户
		err = ModelUser.CreateUser(username, password)
		if err != nil {
			return response.RetError(err, -1)
		}

		response.SetMsg("注册成功")
		return response.RetSuccess("注册成功")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}
