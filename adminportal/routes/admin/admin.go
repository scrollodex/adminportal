package admin

import (
	"fmt"
	"net/http"

	"github.com/scrollodex/scrollodex/adminportal/app"
	"github.com/scrollodex/scrollodex/adminportal/routes/templates"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("SESSION.VALUES = %+v\n", session.Values["profile"])
	fmt.Printf("SESSION.USER = %+v\n", session.Values["user_id"])
	templates.RenderTemplate(w, "admin", session.Values["profile"])
}
