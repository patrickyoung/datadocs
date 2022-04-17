package health

import "github.com/gin-gonic/gin"

func LoadRoutes(app *gin.Engine) {
	health := app.Group("/health")
	{
		health.GET("/ping", ping)
		health.GET("/status", status)
	}

}
