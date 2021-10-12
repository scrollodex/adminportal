package user

import (
	"net/http"

	"github.com/scrollodex/scrollodex/adminportal/app"
	"github.com/scrollodex/scrollodex/adminportal/routes/templates"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "user", session.Values["profile"])
}
