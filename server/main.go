package main

import (
	"htmx-go-todo/todo"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Set the renderer to our HTML template renderer
	e.Renderer = newHTMLRenderer()

	// Root handler shows the home view
	e.GET("/", func(c echo.Context) error {
		// Render the main index template with no parameters, defaulting to the home view
		return c.Render(http.StatusOK, "index", nil)
	})

	// Core view handlers
	AddViewHandlers(e)

	// Register the todo handlers
	todo.AddHandlers(e)

	// Serve static files and log requests
	e.Static("/static", "static")

	// Simple and easy to read request logging
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} (${status})\n",
	}))

	// Accept PORT environment variable, or default to 4000
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Start the server!
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
