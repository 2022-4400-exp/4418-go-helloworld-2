package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/hello", Hello())
    e.GET("/api/hello", ApiHelloGet())

    e.Start(":8080")
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.String(http.StatusOK, "学生番号:4418: 体育会系の社会人１年目です．好きな食べ物はわたあめです． ver.2")
    }
}

func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "4418", "message": "体育会系の社会人１年目です．好きな食べ物はわたあめです．"})
    }
}
