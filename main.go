package main

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neontowel/goyap/pkg/api"
)

//go:embed ui/dist/***
var staticFiles embed.FS

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// API routes
	e.GET("/api/hello", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello from Go!"})
	})

	// Integrate the chat API endpoint
	e.POST("/api/chat", api.ChatHandler) // Use the handler from the api package

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Root:       "ui/dist",
		Filesystem: http.FS(staticFiles),
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
