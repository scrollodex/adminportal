package editrow

import (
	"app"
	"log"
	"strconv"
	"templates"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scrollodex/scrollodex/reslist"
)

/*

This page has many states:

1. GET: Display the HTML form.
2. POST: save the data to reslist (upsert), redirect here?

*/

func EditrowHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{}
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

	site, table, idStr := vars["site"], vars["table"], vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	data["site"] = site
	data["table"] = table
	data["id"] = id

	if site != "bi" && site != "poly" {
		http.Error(w, fmt.Sprintf("No such site: %q", site), http.StatusInternalServerError)
	}

	if table != "categories" && table != "entries" && table != "locations" {
		http.Error(w, fmt.Sprintf("No such table: %q", table), http.StatusInternalServerError)
	}

	dbh, err := reslist.New(fmt.Sprintf("/Users/tlimoncelli/gitthings/scrollodex-db-%s", site))
	if err != nil {
		http.Error(w, fmt.Sprintf("Reslist failed: %q", site), http.StatusInternalServerError)
	}

	switch table {
	case "category":
		d, err := dbh.CategoryGet(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("CategoryGet(%d) failed: %w", id, err), http.StatusInternalServerError)
		}
		data["item"] = d
	case "location":
		d, err := dbh.LocationGet(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("LocationGet(%d) failed: %w", id, err), http.StatusInternalServerError)
		}
		data["item"] = d
	case "entry":
		d, err := dbh.EntryGet(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("EntryGet(%d) failed: %w", id, err), http.StatusInternalServerError)
		}
		data["item"] = d
	}

	templates.RenderTemplate(w, "editrow", data)
}
