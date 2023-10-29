package main

import (
	"htmx-go-todo/todo"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Set the renderer to our HTML template renderer
	e.Renderer = getHTMLRenderer()

	// Views are sub sections of the application, like home or todos
	e.GET("/view/:viewName", ViewHandler)

	// This route allows for URL linking to a specific view
	e.GET("/p/:view", func(c echo.Context) error {
		view := c.Param("view")

		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"View": view,
		})
	})

	// Root handler shows the home view
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"View": "home",
		})
	})

	// Register the todo handlers
	todo.AddHandlers(e)

	// Serve static files and log requests
	e.Static("/static", "static")

	// Simple and easy to read request logging
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} (${status})\n",
	}))

	// Start the server!
	e.Logger.Fatal(e.Start(":4000"))
}
