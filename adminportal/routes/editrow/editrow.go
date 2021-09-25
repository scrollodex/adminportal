package editrow

import (
	"app"
	"templates"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func EditrowHandler(w http.ResponseWriter, r *http.Request) {

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

	site, table, id := vars["site"], vars["table"], vars["id"]

	data["site"] = site
	data["table"] = table
	data["id"] = id

	if site != "bi" && site != "poly" {
		http.Error(w, fmt.Sprintf("No such site: %q", site), http.StatusInternalServerError)
	}

	if table != "categories" && table != "entries" && table != "locations" {
		http.Error(w, fmt.Sprintf("No such table: %q", table), http.StatusInternalServerError)
	}

	//	dbh, err := reslist.New(fmt.Sprintf("/Users/tlimoncelli/gitthings/scrollodex-db-%s", site))
	//	if err != nil {
	//		http.Error(w, fmt.Sprintf("Reslist failed: %q", site), http.StatusInternalServerError)
	//	}
	//	s, err := genCatList(dbh)
	//	if err != nil {
	//		http.Error(w, fmt.Sprintf("No such table: %q", "cat"), http.StatusInternalServerError)
	//	}
	//	data["catlist"] = s
	//	s, err = genLocList(dbh)
	//	if err != nil {
	//		http.Error(w, fmt.Sprintf("No such table: %q", "loc"), http.StatusInternalServerError)
	//	}
	//	data["loclist"] = s
	//	s, err = genStatusList()
	//	if err != nil {
	//		http.Error(w, fmt.Sprintf("status list: %q", "status"), http.StatusInternalServerError)
	//	}
	//	data["statuslist"] = s

	templates.RenderTemplate(w, "editrow", data)
}
