package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mehanizm/airtable"
	"github.com/scrollodex/adminportal/dex/dexmodels"
	"github.com/scrollodex/adminportal/dex/dextidy"
)

func extractLocs(a []*airtable.Record) []dexmodels.LocationYAML {
	b := make([]dexmodels.LocationYAML, len(a))
	for i := range a {
		b[i] = tidyLoc(a[i])
	}
	return b
}

func tidyLoc(a *airtable.Record) dexmodels.LocationYAML {
	f := a.Fields
	//fmt.Printf("F = %+v\n", f)

	b := dexmodels.LocationYAML{
		ID:          int(f["x-LocationID"].(float64)),
		Display:     f["Location"].(string),
		CountryCode: f["x-CountryCode"].(string),
		Region:      f["x-Region"].(string),
	}
	return b
}

func sortLocs(l *[]dexmodels.LocationYAML) {
	sort.Slice((*l), func(i, j int) bool {
		cci := strings.ToLower((*l)[i].CountryCode)
		ccj := strings.ToLower((*l)[j].CountryCode)
		if cci != ccj {
			return cci < ccj
		}
		dni := strings.ToLower((*l)[i].Display)
		dnj := strings.ToLower((*l)[j].Display)
		return dni < dnj
	})
}

func extractCats(a []*airtable.Record) []dexmodels.CategoryYAML {
	b := make([]dexmodels.CategoryYAML, len(a))
	for i := range a {
		b[i] = tidyCat(a[i])
	}
	return b
}

func tidyCat(a *airtable.Record) dexmodels.CategoryYAML {
	f := a.Fields
	//fmt.Printf("F = %+v\n", f)

	b := dexmodels.CategoryYAML{
		Category: dexmodels.Category{
			ID:          int(f["x-CategoryID"].(float64)),
			Name:        f["Name"].(string),
			Description: f["Description"].(string),
			Priority:    int(f["x-Priority"].(float64)),
			Icon:        getString(f, "IconFilename"),
		},
	}
	return b
}

func sortCats(l *[]dexmodels.CategoryYAML) {
	sort.Slice((*l), func(i, j int) bool {
		pi := (*l)[i].Priority
		pj := (*l)[j].Priority
		if pi != pj {
			return pi < pj
		}
		ci := strings.ToLower((*l)[i].Name)
		cj := strings.ToLower((*l)[j].Name)
		return ci < cj
	})
}

func extractEnts(a []*airtable.Record) []dexmodels.PathAndEntry {
	var b []dexmodels.PathAndEntry
	for _, j := range a {
		f := tidyEnt(j)
		f.Title = dextidy.MakeTitle(f)
		path := dextidy.MakePath(f)

		//if j.Status != 1 || j.CategoryID == 0 {
		//// TODO(tlim): Output why this is skipped
		//continue
		//}
		b = append(b, dexmodels.PathAndEntry{
			Path:   path,
			Fields: f,
		})
	}
	return b
}

func getString(f map[string]interface{}, k string) string {
	r := ""
	switch f[k].(type) {
	case string:
		r = f[k].(string)
	default:
	}
	return r
}

func tidyEnt(a *airtable.Record) (b dexmodels.EntryFields) {
	f := a.Fields
	fmt.Printf("F = %+v\n", f)

	b = dexmodels.EntryFields{
		EntryCommon: dexmodels.EntryCommon{
			ID:          int(f["EntryID"].(float64)),
			Salutation:  getString(f, "Salutation"),
			Firstname:   getString(f, "First"),
			Lastname:    getString(f, "Last"),
			Credentials: getString(f, "Suffix"),
			JobTitle:    getString(f, "Job_Title"),
			Company:     getString(f, "Company"),
			ShortDesc:   getString(f, "Short Description"),
			Phone:       getString(f, "Phone"),
			Fax:         getString(f, "Fax"),
			Address:     getString(f, "Address"),
			Email:       getString(f, "Email"),
			Email2:      getString(f, "Email2"),
			Website:     getString(f, "Website"),
			Website2:    getString(f, "Website2"),
			Fees:        getString(f, "Fees"),
			Description: getString(f, "Description"),
			//
		},
		Category:        f["Category"].(string),
		LocationDisplay: f["Location"].(string),
		LastEditDate:    getString(f, "x-lastUpdate"),
	}
	cparts := strings.SplitN(b.LocationDisplay, "-", 2)
	b.Country = cparts[0]

	reg := "unknown"
	if len(cparts) == 1 {
		reg = ""
	} else {
		rparts := strings.SplitN(cparts[1], " ", 2)
		reg = rparts[0]
	}
	//if b.Country == "AT" && reg == "All" {
	//reg = "All"
	//panic("")
	//}
	b.Region = reg

	return b
}

func sortEnts(l *[]dexmodels.PathAndEntry) {
	sort.Slice((*l), func(i, j int) bool {
		pi := (*l)[i].Fields.ID
		pj := (*l)[j].Fields.ID
		return pi < pj
	})
}
