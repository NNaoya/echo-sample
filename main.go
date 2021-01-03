package main

import (
    "net/http"
    "github.com/myapp/conf"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/myapp/api"
    _ "github.com/myapp/docs"
    echoSwagger "github.com/swaggo/echo-swagger"
    "flag"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:1313
// @BasePath /
func main() {
    e := echo.New()

    // 設定ファイル読込
    setConfig()

    // 確認用出力
    c := config.Config.Database
    fmt.Printf("DBユーザー::%s", c.User)

    // Routes
    v1 := e.Group("/api/v1")
    {
        v1.GET("/", hello)
        v1.POST("/login", api.Login())
        v1.GET("/swagger/*", echoSwagger.WrapHandler)
    }
    r := e.Group("/api/v1/restricted")
    {
        r.Use(middleware.JWT([]byte("secret")))
        r.GET("/welcome", api.Restricted())
    }
    e.Logger.Fatal(e.Start(":1323"))
}

// @Summary get string
// @Description get "Hello, World!" of string
// @Accept plain
// @Produce plain
// @Success 200
// @Router / [get]
func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}

// getUsers is getting users.
// @Summary get users
// @Description get users in a group
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user [get]
func show(c echo.Context) error {
    u := new(User)
    return c.JSON(http.StatusOK, u)
}

// コマンドライン引数でproが指定された場合は本番環境の設定を取得
func setConfig() {
  env := "development"
  flag.Parse()
  if args := flag.Args(); 0 < len(args) && args[0] == "pro" {
    env = "production"
  }
  config.SetEnvironment(env)
}