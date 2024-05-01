package echotmpl

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func createServer(config Config) *echo.Echo {
	var app = echo.New()
	app.Renderer = GetRenderer(config)
	return app
}

func TestGetRenderer(t *testing.T) {
	var renderer = GetRenderer(Config{
		Root:          "test/views",
		Extension:     ".html",
		LayoutFolder:  "_layouts",
		DefaultLayout: "base",
	})

	if renderer == nil {
		t.FailNow()
	}
}

func TestRootTemplate(t *testing.T) {
	app := createServer(Config{
		Root:          "test/views",
		Extension:     ".html",
		LayoutFolder:  "_layouts",
		DefaultLayout: "base",
	})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := app.NewContext(req, res)

	c.Render(http.StatusOK, "index.html", nil)

	if !strings.Contains(res.Body.String(), "Base Template") {
		t.FailNow()
	}

	if !strings.Contains(res.Body.String(), "Hello World") {
		t.FailNow()
	}
}

func TestSubfolderTemplate(t *testing.T) {
	app := createServer(Config{
		Root:          "test/views",
		Extension:     ".html",
		LayoutFolder:  "_layouts",
		DefaultLayout: "base",
	})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := app.NewContext(req, res)

	c.Render(http.StatusOK, "users/index.html", nil)

	if !strings.Contains(res.Body.String(), "Users") {
		t.FailNow()
	}
}

func TestChangeLayout(t *testing.T) {
	app := createServer(Config{
		Root:          "test/views",
		Extension:     ".html",
		LayoutFolder:  "_layouts",
		DefaultLayout: "base",
	})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := app.NewContext(req, res)

	c.Set("layout", "admin")
	c.Render(http.StatusOK, "users/index.html", nil)

	if !strings.Contains(res.Body.String(), "Admin Template") {
		t.FailNow()
	}

	if !strings.Contains(res.Body.String(), "Users") {
		t.FailNow()
	}
}

func TestWrongTemplate(t *testing.T) {
	app := createServer(Config{
		Root:          "test/views",
		Extension:     ".html",
		LayoutFolder:  "_layouts",
		DefaultLayout: "base",
	})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := app.NewContext(req, res)

	err := c.Render(http.StatusOK, "invalid.html", nil)

	if err == nil {
		t.FailNow()
	}
}
