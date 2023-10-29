package main

import (
	"htmx-go-todo/todo"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Renderer = getHTMLRenderer()

	e.GET("/view/:viewName", ViewHandler)

	e.GET("/p/:view", func(c echo.Context) error {
		view := c.Param("view")
		log.Println("#### view:", view)

		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"View": view,
		})
	})

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"View": "home",
		})
	})

	todo.AddHandlers(e)

	e.Static("/static", "static")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} (${status})\n",
	}))

	e.Logger.Fatal(e.Start(":4000"))
}
