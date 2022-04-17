package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patrickyoung/datadocs/templates"
	"github.com/patrickyoung/datadocs/utils"
)

func TestPingStatusOK(t *testing.T) {
	app := utils.GetApp()

	LoadRoutes(app)

	req, _ := http.NewRequest("GET", "/health/ping", nil)

	utils.TestHTTPResponse(t, app, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		return statusOK
	})
}

func TestStatusPageOK(t *testing.T) {
	app := utils.GetApp()

	templates.LoadTemplates(app)
	LoadRoutes(app)

	req, _ := http.NewRequest("GET", "/health/status", nil)

	utils.TestHTTPResponse(t, app, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		return statusOK
	})
}
