package zingdata

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scrollodex/adminportal/app"
	"github.com/scrollodex/dex/reslist"
)

func ZingDataHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if profile, ok := session.Values["profile"]; ok {
		mp, ok := profile.(map[string]interface{})
		if !ok {
			panic("An entire interface changed type. I just can't.")
		}
		_ = mp
	}

	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	site, table := vars["site"], vars["table"]
	if site != "bi" && site != "poly" {
		http.Error(w, fmt.Sprintf("No such site: %q", site), http.StatusInternalServerError)
	}

	//dbh, err := reslist.New(fmt.Sprintf("git@scrollodex-github.com:scrollodex/scrollodex-db-%s.git", site))
	dbh, err := reslist.New(fmt.Sprintf("/Users/tlimoncelli/gitthings/scrollodex-db-%s", site))
	//dbh, err := reslist.New(os.GetEnv("ADMINPORTAL_RESBASE", site)
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
		b, err := json.MarshalIndent(l, "", "\t")
		if err != nil {
			http.Error(w, fmt.Sprintf("CatJSON failed: %s", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	case "locations":
		l, err := dbh.LocationList()
		if err != nil {
			http.Error(w, fmt.Sprintf("LocList failed: %s", err), http.StatusInternalServerError)
		}
		b, err := json.MarshalIndent(l, "", "\t")
		if err != nil {
			http.Error(w, fmt.Sprintf("LocJSON failed: %s", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	case "entries":
		l, err := dbh.EntryList()
		if err != nil {
			http.Error(w, fmt.Sprintf("EntList failed: %s", err), http.StatusInternalServerError)
		}
		b, err := json.MarshalIndent(l, "", "\t")
		if err != nil {
			http.Error(w, fmt.Sprintf("EntJSON failed: %s", err), http.StatusInternalServerError)
		}
		zingdata = string(b)

	default:
		http.Error(w, fmt.Sprintf("No such table: %q", table), http.StatusInternalServerError)
	}

	w.Write([]byte(zingdata))
}
