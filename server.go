package main

import (
	"editether/database"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

const PORT = ":8080"

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	database.InitDb("./editether.db")

	e := echo.New()
	e.Static("/static", "public")
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.Logger.Fatal(e.Start(PORT))
}
