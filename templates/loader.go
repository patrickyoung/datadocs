package templates

import (
	"embed"
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed *
var f embed.FS

func LoadTemplates(r *gin.Engine) {
	myFuncMap := template.FuncMap{
		"formatAsDate": formatAsDate,
	}

	// Load templates from embedded filesystem
	templ := template.Must(template.New("").Funcs(myFuncMap).ParseFS(f, "**/*.tmpl"))
	r.SetHTMLTemplate(templ)
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
