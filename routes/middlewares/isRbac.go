package middlewares

import (
	"net/http"

	"github.com/scrollodex/adminportal/app"
	"github.com/scrollodex/adminportal/rbac"
)

// IsRbacEditor is middleware that requires "Editor" entitlement.
func IsRbacEditor(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if profile, ok := session.Values["profile"]; !ok {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
	} else {
		//fmt.Printf("PROFILE = %+v\n", profile)
		//fmt.Printf("PROFILE TYPE = %T\n", profile)

		username := rbac.MakeUsername(profile)
		//fmt.Printf("USERNAME = %s\n", username)
		if !rbac.Can(username, "editor") {
			http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		}
		next(w, r)
	}
}
