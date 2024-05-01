# Echo Tmpl

Package to add the possibility to use subfolders for your views

Why should I use it: HTMX apps

## Install

`go get github.com/astlaure/echo-tmpl`

## Usage

```go
var app = echo.New()
app.Renderer = echotmpl.GetRenderer(echotmpl.Config{
    Views: "resources/views",
    LayoutFolder: "_layouts",
    DefaultLayout: "base",
})
```
