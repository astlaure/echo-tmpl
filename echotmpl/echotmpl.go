package echotmpl

import (
	"errors"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates     map[string]*template.Template
	defaultLayout string
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		return errors.New("Template not found -> " + name)
	}

	return tmpl.ExecuteTemplate(w, selectLayout(c, t.defaultLayout), data)
}

func registerTemplates(config Config) map[string]*template.Template {
	var elements = make(map[string]*template.Template)
	var layouts []string

	// Load the layouts
	filepath.WalkDir(config.Root+"/"+config.LayoutFolder, func(path string, d fs.DirEntry, err error) error {
		path = strings.ReplaceAll(path, "\\", "/") // Windows path fix

		if !d.IsDir() && strings.HasSuffix(path, config.Extension) {
			layouts = append(layouts, path)
		}

		return nil
	})

	// Load the templates
	filepath.WalkDir(config.Root, func(path string, d fs.DirEntry, err error) error {
		path = strings.ReplaceAll(path, "\\", "/") // Windows path fix

		if !d.IsDir() && strings.HasSuffix(path, config.Extension) {
			var name, _ = strings.CutPrefix(
				strings.ReplaceAll(path, config.Root, ""),
				"/",
			)

			var files = append(layouts, path)
			elements[name] = template.Must(template.ParseFiles(files...))
		}

		return nil
	})

	return elements
}

func selectLayout(c echo.Context, defaultLayout string) string {
	layout, ok := c.Get("layout").(string)

	if !ok {
		return defaultLayout
	}

	return layout
}

type Config struct {
	Root          string
	Extension     string
	LayoutFolder  string
	DefaultLayout string
}

func GetRenderer(config Config) *Template {
	return &Template{
		templates:     registerTemplates(config),
		defaultLayout: config.DefaultLayout,
	}
}
