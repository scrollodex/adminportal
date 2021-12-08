package edit

import (
	"fmt"
	"net/http"
	"os"

	"github.com/scrollodex/adminportal/app"
	"github.com/scrollodex/adminportal/routes/templates"
	"github.com/scrollodex/dex/dextidy"
	"github.com/scrollodex/dex/reslist"

	"github.com/gorilla/mux"
)

type nameVal = struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func EditHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]string{}
	data["nickname"] = "friend"
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

	if table != "categories" && table != "entries" && table != "locations" {
		http.Error(w, fmt.Sprintf("No such table: %q", table), http.StatusInternalServerError)
	}

	if table == "entries" {
		//dbh, err := reslist.New(fmt.Sprintf("/Users/tlimoncelli/gitthings/scrollodex-db-%s", site))
		dbh, err := reslist.New(os.Getenv("ADMINPORTAL_DB_CONNECTSTRING"), site)
		if err != nil {
			http.Error(w, fmt.Sprintf("Reslist failed: %q", site), http.StatusInternalServerError)
		}

		s, err := dextidy.GenCatList(dbh)
		if err != nil {
			http.Error(w, fmt.Sprintf("No such table: %q", "cat"), http.StatusInternalServerError)
		}
		data["catlist"] = s

		s, err = dextidy.GenLocList(dbh)
		if err != nil {
			http.Error(w, fmt.Sprintf("No such table: %q", "loc"), http.StatusInternalServerError)
		}
		data["loclist"] = s

		s, err = dextidy.GenStatusList()
		if err != nil {
			http.Error(w, fmt.Sprintf("status list: %q", "status"), http.StatusInternalServerError)
		}
		data["statuslist"] = s
	}

	templates.RenderTemplate(w, "edit", data)
}
