package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

var redirects = map[string]string{
	"discord": "https://discord.com/invite/AjsxPxdKkB",
	"steam": "https://store.steampowered.com/app/312530/Duck_Game/",
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func setupRenderer(app *echo.Echo) {
	app.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
}

func setupErrorHandler(app *echo.Echo) {
	app.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		errorPage := fmt.Sprintf("%d.html", code)
		if err := c.File(errorPage); err != nil {
			c.Logger().Error(err)
		}
		c.Logger().Error(err)
	}
}

func main() {
	app := echo.New()

	setupRenderer(app)
	setupErrorHandler(app)

	app.GET("/", Home)
	app.Static("/static", "static")
	app.GET("/:path", Redirect)

	app.Logger.Fatal(app.Start(":3825"))
}

func Home(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "home.html", "")
}

func Redirect(ctx echo.Context) error {
	path := ctx.Param("path")
	url, exists := redirects[path]
	if exists {
		return ctx.Redirect(http.StatusPermanentRedirect, url)
	} else {
		return ctx.NoContent(http.StatusNotFound)
	}
}