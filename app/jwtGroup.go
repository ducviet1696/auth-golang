package api

import (
    "authGoLang/app/handlers"

    "github.com/labstack/echo"
)

func JwtGroup(g *echo.Group) {
    g.GET("/main", handlers.MainJwt)
}
