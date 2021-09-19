package edit

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app"
	"templates"

	"github.com/gorilla/mux"
	"github.com/scrollodex/scrollodex/reslist"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]string{}
	if profile, ok := session.Values["profile"]; ok {
		mp, ok := profile.(map[string]interface{})
		if !ok {
			panic("An entire interface changed type. I just can't.")
		}
		data["nickname"] = mp["nickname"].(string)
	}

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	site, table := vars["site"], vars["table"]

	data["site"] = site
	data["table"] = table

	if site != "bi" && site != "poly" {
		http.Error(w, fmt.Sprintf("No such site: %q", site), http.StatusInternalServerError)
	}

	dbh, err := reslist.New(fmt.Sprintf("git@scrollodex-github.com:scrollodex/scrollodex-db-%s.git", site))
	if err != nil {
		http.Error(w, fmt.Sprintf("CatList failed: %s", err), http.StatusInternalServerError)
	}
	var zingdata string
	switch table {

	case "categories":
		l, err := dbh.CategoryList()
		if err != nil {
			http.Error(w, fmt.Sprintf("CatList failed: %s", err), http.StatusInternalServerError)
		}
		b, err := json.Marshal(l)
		if err != nil {
			http.Error(w, fmt.Sprintf("CatJSON failed: %s", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	case "locations":
		l, err := dbh.LocationList()
		if err != nil {
			http.Error(w, fmt.Sprintf("LocList failed: %s", err), http.StatusInternalServerError)
		}
		b, err := json.Marshal(l)
		if err != nil {
			http.Error(w, fmt.Sprintf("LocJSON failed: %s", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	case "entries":
		l, err := dbh.EntryList()
		if err != nil {
			http.Error(w, fmt.Sprintf("EntList failed: %s", err), http.StatusInternalServerError)
		}
		b, err := json.Marshal(l)
		if err != nil {
			http.Error(w, fmt.Sprintf("EntJSON failed: %s", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	default:
		http.Error(w, fmt.Sprintf("No such table: %q", table), http.StatusInternalServerError)
	}

	data["data"] = zingdata

	templates.RenderTemplate(w, "edit", data)
}
