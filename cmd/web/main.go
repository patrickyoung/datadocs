package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/patrickyoung/datadocs/health"
	"github.com/patrickyoung/datadocs/templates"
)

type WebServerConfig struct {
	TrustedProxies []string
	Port           int32
	URL            string
}

func main() {
	initWebApp(WebServerConfig{
		TrustedProxies: []string{"127.0.0.1"},
		Port:           8080,
		URL:            "http://localhost:8080",
	})
}

func initWebApp(config WebServerConfig) {
	app := gin.Default()
	app.SetTrustedProxies(config.TrustedProxies)

	templates.LoadTemplates(app)
	health.LoadRoutes(app)

	app.Run(fmt.Sprintf(":%d", config.Port))
}
