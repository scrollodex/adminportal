package editrow

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/scrollodex/adminportal/app"
	"github.com/scrollodex/adminportal/dex/dextidy"
	"github.com/scrollodex/adminportal/dex/reslist"
	"github.com/scrollodex/adminportal/routes/templates"
)

/*

This page has many states:

1. GET: Display the HTML form.
2. POST: save the data to reslist (upsert), redirect here?

*/

// Handler renders the page.
func Handler(w http.ResponseWriter, r *http.Request) {

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

	//dbh, err := reslist.New(fmt.Sprintf("/Users/tlimoncelli/gitthings/scrollodex-db-%s", site))
	dbh, err := reslist.New(os.Getenv("ADMINPORTAL_DB_CONNECTSTRING"), site)
	if err != nil {
		http.Error(w, fmt.Sprintf("Reslist failed: %q", site), http.StatusInternalServerError)
	}

	switch table {

	case "categories":
		d, err := dbh.CategoryGet(id)
		fmt.Fprintf(os.Stderr, "DEBUG: err=%v d=%+v\n", err, d)
		if err != nil {
			http.Error(w, fmt.Sprintf("CategoryGet(%d) failed: %s", id, err), http.StatusInternalServerError)
		}
		data["item"] = d

	case "locations":
		d, err := dbh.LocationGet(id)
		fmt.Fprintf(os.Stderr, "DEBUG: err=%v d=%+v\n", err, d)
		if err != nil {
			http.Error(w, fmt.Sprintf("LocationGet(%d) failed: %s", id, err), http.StatusInternalServerError)
		}
		data["item"] = d

	case "entries":
		d, err := dbh.EntryGet(id)
		fmt.Fprintf(os.Stderr, "DEBUG: err=%v d=%+v\n", err, d)
		if err != nil {
			http.Error(w, fmt.Sprintf("EntryGet(%d) failed: %s", id, err), http.StatusInternalServerError)
		}
		data["item"] = d

		if s, err := dextidy.CatNameVal(dbh); err != nil {
			http.Error(w, fmt.Sprintf("NameVal(cat) err: %s", err), http.StatusInternalServerError)
		} else {
			data["catnvl"] = s
		}

		if s, err := dextidy.LocNameVal(dbh); err != nil {
			http.Error(w, fmt.Sprintf("NameVal(loc) err: %s", err), http.StatusInternalServerError)
		} else {
			data["locnvl"] = s
		}

	}

	templates.RenderTemplate(w, "editrow", data)
}
