package edit

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app"
	"templates"

	"github.com/gorilla/mux"
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

	//dbh := reslist.New(ft.Sprintf("git@scrollodex-github.com:scrollodex/scrollodex-db-%s.git", site))
	var dbh interface{}
	var zingdata string
	switch table {

	case "categories":
		l := dbh.CategoryList()
		if err != nil {
			http.Error(w, fmt.Errorf("CatList failed: %w", err), http.StatusInternalServerError)
		}
		b, err := json.MarshalIndent(l)
		if err != nil {
			http.Error(w, fmt.Errorf("CatJSON failed: %w", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	case "locations":
		l := dbh.LocationList()
		if err != nil {
			http.Error(w, fmt.Errorf("LocList failed: %w", err), http.StatusInternalServerError)
		}
		b, err := json.MarshalIndent(l)
		if err != nil {
			http.Error(w, fmt.Errorf("LocJSON failed: %w", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	case "entries":
		l := dbh.EntryList()
		if err != nil {
			http.Error(w, fmt.Errorf("EntList failed: %w", err), http.StatusInternalServerError)
		}
		b, err := json.MarshalIndent(l)
		if err != nil {
			http.Error(w, fmt.Errorf("EntJSON failed: %w", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	default:
		http.Error(w, fmt.Errorf("No such table: %q", table), http.StatusInternalServerError)
	}

	data["data"] = zingdata

	templates.RenderTemplate(w, "edit", data)
}
