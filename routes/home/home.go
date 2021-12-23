package home

import (
	"net/http"

	"github.com/scrollodex/adminportal/routes/templates"
)

// Handler renders the page.
func Handler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}
