package edit

import (
	"app"
	"templates"

	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/scrollodex/scrollodex/dextidy"
	"github.com/scrollodex/scrollodex/reslist"

	"github.com/gorilla/mux"
)

type nameVal = struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func genCatList(dbh reslist.Databaser) (string, error) {
	orig, err := dbh.CategoryList()
	if err != nil {
		return "", err
	}
	var nvl []*nameVal
	for _, item := range orig {
		n := &nameVal{Name: item.Name, Value: item.ID}
		nvl = append(nvl, n)
	}
	sort.Slice(nvl, func(i, j int) bool { return nvl[i].Name < nvl[j].Name })
	b, err := json.MarshalIndent(&nvl, "", "\t")
	if err != nil {
		return "", err
	}
	s := string(b)
	return s, nil
}

func genLocList(dbh reslist.Databaser) (string, error) {
	orig, err := dbh.LocationList()
	if err != nil {
		return "", err
	}
	var nvl []nameVal
	for _, item := range orig {
		n := nameVal{Name: dextidy.MakeDisplayLoc(item), Value: item.ID}
		nvl = append(nvl, n)
	}
	sort.Slice(nvl, func(i, j int) bool {
		// Sort the "-All" of each country to the top.
		a := nvl[i].Name
		b := nvl[j].Name
		if a[:3] == b[:3] {
			if strings.Contains(a, "-All") {
				return true
			}
			if strings.Contains(b, "-All") {
				return false
			}
		}
		//  Otherwise, sort lexigraphically.
		return nvl[i].Name < nvl[j].Name
	})

	b, err := json.MarshalIndent(&nvl, "", "\t")
	if err != nil {
		return "", err
	}
	s := string(b)
	return s, nil
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
	data["loclist"] = `[]`
	dbh, err := reslist.New(fmt.Sprintf("/Users/tlimoncelli/gitthings/scrollodex-db-%s", site))
	if err != nil {
		http.Error(w, fmt.Sprintf("Reslist failed: %q", site), http.StatusInternalServerError)
	}

	s, err := genCatList(dbh)
	if err != nil {
		http.Error(w, fmt.Sprintf("No such table: %q", "cat"), http.StatusInternalServerError)
	}
	data["catlist"] = s

	s, err = genLocList(dbh)
	if err != nil {
		http.Error(w, fmt.Sprintf("No such table: %q", "loc"), http.StatusInternalServerError)
	}
	data["loclist"] = s

	templates.RenderTemplate(w, "edit", data)
}
