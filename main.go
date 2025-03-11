package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neontowel/goyap/pkg/api"
)

//go:embed ui/dist/***
var staticFiles embed.FS

func main() {
	// Load the environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: No .env file found")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure CORS middleware globally, applicable to all routes
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{echo.POST, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	e.GET("/api/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello from Go!"})
	})

	e.POST("/api/chat", api.ChatHandler)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Root:       "ui/dist",
		Filesystem: http.FS(staticFiles),
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
