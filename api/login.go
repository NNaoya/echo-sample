package api

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "time"
    "github.com/dgrijalva/jwt-go"
)

func Login() echo.HandlerFunc {
    return func(c echo.Context) error {
        userName := c.FormValue("username")
        password := c.FormValue("password")

        if userName == "test" && password == "test" {
            token := jwt.New(jwt.SigningMethodHS256)

            claims := token.Claims.(jwt.MapClaims)
            claims["name"] = "test"
            claims["admin"] = true
            claims["iat"] = time.Now().Unix()
            claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

            // generate encoded token and send it as response
            t, err := token.SignedString([]byte("secret"))
            if err != nil {
                return err
            }
            return c.JSON(http.StatusOK, map[string]string{
                "token": t,
            })
        }

        return echo.ErrUnauthorized
    }
}

func Restricted() echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        name := claims["name"].(string)

        return c.String(http.StatusOK, "Welcome" + name + "!")
    }
}