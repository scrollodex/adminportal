package main

import (
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
			Icon:        f["IconFilename"].(string),
		},
	}
	return b
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
	//fmt.Printf("F = %+v\n", f)

	b = dexmodels.EntryFields{
		//Title: f[""].(string),

		EntryCommon: dexmodels.EntryCommon{
			ID:          int(f["EntryID"].(float64)),
			Salutation:  getString(f, "salutation"),
			Firstname:   getString(f, "first_name"),
			Lastname:    getString(f, "last_name"),
			Credentials: getString(f, "credentials"),
			JobTitle:    getString(f, "job_title"),
			Company:     getString(f, "company"),
			ShortDesc:   getString(f, "short_desc"),
			Phone:       getString(f, "phone"),
			Fax:         getString(f, "fax"),
			Address:     getString(f, "address"),
			Email:       getString(f, "email"),
			Email2:      getString(f, "email2"),
			Website:     getString(f, "website"),
			Website2:    getString(f, "website2"),
			Fees:        getString(f, "fees"),
			Description: getString(f, "description"),
		},

		//
		//Category:        f["Category"].(string),
		//LocationDisplay: f["Location"].(string),
		//Country:         f["Country"].(string),
		//Region:          f["Region"].(string),
		//LastEditDate:    f[""].(string),
	}

	return b
}
