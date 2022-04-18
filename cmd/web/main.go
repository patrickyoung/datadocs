package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/patrickyoung/datadocs/health"
	"github.com/patrickyoung/datadocs/templates"
	"github.com/patrickyoung/datadocs/utils"
	"github.com/patrickyoung/datadocs/ydoc"
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
	ydoc.LoadRoutes(app)

	utils.OpenBrowser(fmt.Sprintf("%s/health/ping", config.URL))

	app.Run(fmt.Sprintf(":%d", config.Port))
}
