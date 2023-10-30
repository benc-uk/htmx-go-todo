package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Implement echo.Renderer interface
type HTMLRenderer struct {
	templates *template.Template
}

func (r *HTMLRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

// This function returns a new HTMLRenderer which loads templates from
// the ./templates directory
func getHTMLRenderer() *HTMLRenderer {
	return &HTMLRenderer{
		templates: template.Must(template.ParseGlob("./templates/**/*.html")),
	}
}

// Register the view handlers
func AddViewHandlers(e *echo.Echo) {
	// Views are sub sections of the application, like home or todos
	// These are rendered under the navbar
	e.GET("/view/:viewName", func(c echo.Context) error {
		name := c.Param("viewName")

		err := c.Render(http.StatusOK, "view/"+name, nil)

		if err != nil {
			return c.String(http.StatusNotFound, "Not Found")
		}

		return nil
	})

	// This route allows for URLs linking directly to a specific view
	// e.g. http://localhost:4000/p/list-todos
	e.GET("/p/:viewName", func(c echo.Context) error {
		viewName := c.Param("viewName")

		// Render the main index template with the view name
		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"View": viewName,
		})
	})
}
