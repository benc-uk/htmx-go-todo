package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTMLRenderer struct {
	templates *template.Template
}

func (r *HTMLRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

func ViewHandler(c echo.Context) error {
	name := c.Param("viewName")

	err := c.Render(http.StatusOK, name, nil)

	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return nil
}

func getHTMLRenderer() *HTMLRenderer {
	return &HTMLRenderer{
		templates: template.Must(template.ParseGlob("./templates/**/*.html")),
	}
}
