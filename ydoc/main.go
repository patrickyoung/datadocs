package ydoc

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type DataDoc struct {
	Id        string    `yaml:"id"`
	Version   string    `yaml:"version"`
	Namespace string    `yaml:"namespace"`
	Modified  time.Time `yaml:"modified"`
	Template  string    `yaml:"template"`
	Data      string    `yaml:"content"`
}

func UnmarshalData(datadoc DataDoc) map[string]interface{} {
	var data map[string]interface{}
	err := yaml.Unmarshal([]byte(datadoc.Data), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func getData(c *gin.Context) {
	datadoc := DataDoc{
		Id:        "001",
		Version:   "001-1",
		Namespace: "ydoc",
		Modified:  time.Now(),
		Template:  "ydoc/simple.tmpl",
		Data:      "myvalue: TEST-002-CONTENT",
	}
	data := UnmarshalData(datadoc)

	c.HTML(http.StatusOK, "docs/myvalue.tmpl", data)
}

func LoadRoutes(app *gin.Engine) {
	docs := app.Group("/docs")
	{
		docs.GET("/show", getData)
	}

}
