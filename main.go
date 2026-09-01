// Package main starts the URL shortener HTTP server.
package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/getsentry/sentry-go"
)

func setupRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
	Dsn: os.Getenv("SENTRY_DSN"),

	SendDefaultPII: true,

	EnableTracing:        false,
	TracesSampleRate:     0,
	DisableClientReports: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := setupRouter().Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
