package unauthorized

import (
	"net/http"

	"github.com/scrollodex/adminportal/app"
	"github.com/scrollodex/adminportal/rbac"
	"github.com/scrollodex/adminportal/routes/templates"
)

func UnauthorizedHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var vars = map[string]string{}
	vars["username"] = "NOT A USER"
	if profile, ok := session.Values["profile"]; ok {
		vars["username"] = rbac.MakeUsername(profile)
	}
	// TODO(tlim): Set nickname to the auth0 nickname.
	vars["nickname"] = "friend"

	templates.RenderTemplate(w, "unauthorized", vars)
}
