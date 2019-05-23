package api

import (
    "authGoLang/app/handlers"
    "github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
    e.GET("/login", handlers.Login)
}
