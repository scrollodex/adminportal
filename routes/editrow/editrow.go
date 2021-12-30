package editrow

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/scrollodex/adminportal/app"
	"github.com/scrollodex/adminportal/dex/dexmodels"
	"github.com/scrollodex/adminportal/dex/dextidy"
	"github.com/scrollodex/adminportal/dex/reslist"
	"github.com/scrollodex/adminportal/rbac"
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

	// Session (cookie) values:
	var data = map[string]interface{}{}

	// Defaults:
	data["nickname"] = "friend"
	emailAddr := "n/a"

	if profile, ok := session.Values["profile"]; ok {
		mp, ok := profile.(map[string]interface{})
		if !ok {
			panic("An entire interface changed type. I just can't.")
		}
		data["nickname"] = mp["nickname"].(string)
		emailAddr = rbac.EmailOf(mp["iss"].(string), mp["sub"].(string))
	}

	fmt.Printf("DEBUG: =================================== NEW REQUEST\n")

	// URL path parameters:

	vars := mux.Vars(r)
	//w.WriteHeader(http.StatusOK)

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

	// URL params:

	//param := r.URL.Query()
	//fmt.Printf("DEBUG: PARAM = %+v\n", param)
	err = r.ParseForm()
	if err != nil {
		// in case of any error
		http.Error(w, fmt.Sprintf("ParseForm failed: %s", err), http.StatusInternalServerError)
		return
	}
	//value := r.Form.Get("parameter_name")
	fmt.Printf("DEBUG: METHOD = %v\n", r.Method)
	fmt.Printf("DEBUG: PARAMS = %+v\n", r.Form)

	//dbh, err := reslist.New(fmt.Sprintf("/Users/tlimoncelli/gitthings/scrollodex-db-%s", site))
	dbh, err := reslist.New(os.Getenv("ADMINPORTAL_DB_CONNECTSTRING"), site)
	if err != nil {
		http.Error(w, fmt.Sprintf("Reslist failed: %q", site), http.StatusInternalServerError)
	}
	fmt.Printf("DEBUG dbh=%v\n", dbh)

	// At this point any data from the request is gathered and validated.

	// If this is a POST, save the data.
	if r.Method == "POST" {

		switch table {

		case "categories":
			priority, _ := strconv.Atoi(r.Form["text-1632675435086"][0])
			err = dbh.CategoryStore(dexmodels.Category{
				ID:          id,
				Name:        r.Form["text-1632675383103"][0],
				Description: r.Form["text-1632675406583"][0],
				Priority:    priority,
				Icon:        r.Form["text-1632675419960"][0],
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("Store failed: %q", err), http.StatusInternalServerError)
				return
			}

		case "locations":
			err = dbh.LocationStore(dexmodels.Location{
				ID:          id,
				CountryCode: r.Form["text-1632673613322"][0],
				Region:      r.Form["text-1632673614839"][0],
				Comment:     r.Form["text-1632673617487"][0],
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("Store failed: %q", err), http.StatusInternalServerError)
				return
			}

		case "entries":

			categoryID, _ := strconv.Atoi(r.Form["select-1632676732741"][0])
			locationID, _ := strconv.Atoi(r.Form["select-1632676809195"][0])
			err = dbh.EntryStore(dexmodels.Entry{
				EntryCommon: dexmodels.EntryCommon{
					ID:          id,
					Salutation:  r.Form["text-1632676150550"][0],
					Firstname:   r.Form["text-1632676152332"][0],
					Lastname:    r.Form["text-1632676155685"][0],
					Credentials: r.Form["text-1632676386410"][0],
					JobTitle:    r.Form["text-1632676481648"][0],
					Company:     r.Form["text-1632676572979"][0],
					ShortDesc:   r.Form["text-1632676600233"][0],
					Phone:       r.Form["text-1632676662873"][0],
					Fax:         r.Form["text-1632676648914"][0],
					Address:     r.Form["text-1632676654130"][0],
					Email:       r.Form["text-1632676665914"][0],
					Email2:      r.Form["text-1632676669762"][0],
					Website:     r.Form["text-1632676678154"][0],
					Website2:    r.Form["text-1632676674523"][0],
					Fees:        r.Form["text-1632676682890"][0],
					Description: r.Form["textarea-1632676686169"][0],
				},
				// Form data
				PrivateAdminNotes:   r.Form["textarea-1632677017043"][0],
				PrivateContactEmail: r.Form["text-1632677014693"][0],

				// Indexes:
				CategoryID: categoryID,
				LocationID: locationID,

				// Calculated
				//LastEditDate:      time.Now().Format(time.RFC3339),
				LastEditDate:      time.Now().UTC().Format("2006-01-02 15:04 UTC"),
				PrivateLastEditBy: emailAddr,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("Store failed: %q", err), http.StatusInternalServerError)
				return
			}

		}
	}

	// Generate the new web page:

	switch table {

	case "categories":
		d, err := dbh.CategoryGet(id)
		fmt.Fprintf(os.Stderr, "DEBUG: err=%v d=%+v\n", err, d)
		if err != nil {
			http.Error(w, fmt.Sprintf("CategoryGet(%d) failed: %s", id, err), http.StatusInternalServerError)
		}
		data["item"] = d

	case "locations":
		fmt.Printf("DEBUG: id=%v %T\n", id, id)
		d, err := dbh.LocationGet(id)
		fmt.Fprintf(os.Stderr, "DEBUG: err=%v d=%+v\n", err, d)
		if err != nil {
			http.Error(w, fmt.Sprintf("LocationGet(%d) failed: %s", id, err), http.StatusInternalServerError)
		}
		data["item"] = d

	case "entries":
		d, err := dbh.EntryGet(id)
		if d.PrivateLastEditBy == "" {
			d.PrivateLastEditBy = "unknown"
		}
		if d.LastEditDate == "" {
			d.LastEditDate = "unknown"
		}
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
