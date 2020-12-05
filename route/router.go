package route

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/myapp/api"
    _ "github.com/myapp/cmd/myapp/docs"
      echoSwagger "github.com/swaggo/echo-swagger"
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
func Init() *echo.Echo {
    e := echo.New()

    // Set Bundle Middleware

    // Routes
    v1 := e.Group("/api/v1")
    {
        v1.GET("/", hello)
        v1.GET("/ping", func(c echo.Context) error {
          return c.String(http.StatusOK, "Hello, World!")
        })
        v1.GET("/test", func(c echo.Context) error {
           return c.String(http.StatusOK, "Let's study!!")
        })
        v1.GET("/test1", func(c echo.Context) error {
          return c.String(http.StatusOK, "Let's study!!!!")
        })
        v1.GET("/test2", func(c echo.Context) error {
          return c.String(http.StatusOK, "Let's study!!!!")
        })
        v1.GET("/user", show)
        v1.POST("/login", api.Login())
        v1.GET("/swagger/*", echoSwagger.WrapHandler)
    }

    r := e.Group("/api/v1/restricted")
    {
        r.Use(middleware.JWT([]byte("secret")))
        r.GET("/welcome", api.Restricted())
    }

    return e
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