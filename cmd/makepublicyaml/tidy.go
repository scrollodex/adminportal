package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/scrollodex/scrollodex/dexmodels"
	"github.com/scrollodex/scrollodex/reslist"
)

func getCatRaw(dbh reslist.Databaser) ([]dexmodels.Category, error) {
	l, err := dbh.CategoryList()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func tidyCats(a []dexmodels.Category) []dexmodels.CategoryYAML {
	b := make([]dexmodels.CategoryYAML, len(a))
	for i := range a {
		b[i] = tidyCat(&(a[i]))
	}
	return b
}

func tidyCat(a *dexmodels.Category) (b dexmodels.CategoryYAML) {
	b = dexmodels.CategoryYAML{}
	b.ID = a.ID
	b.Name = a.Name
	b.Description = a.Description
	b.Priority = a.Priority
	return b
}

func sortCat(l *[]dexmodels.CategoryYAML) {
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

func getLocRaw(dbh reslist.Databaser) ([]dexmodels.Location, error) {
	l, err := dbh.LocationList()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func tidyLocs(a []dexmodels.Location) []dexmodels.LocationYAML {
	b := make([]dexmodels.LocationYAML, len(a))
	for i := range a {
		b[i] = tidyLoc(&(a[i]))
	}
	return b
}

func tidyLoc(a *dexmodels.Location) dexmodels.LocationYAML {
	b := dexmodels.LocationYAML{}
	b.ID = a.ID
	b.CountryCode = a.CountryCode
	b.Region = a.Region
	if b.CountryCode == "ZZ" {
		if a.Comment == "" {
			b.Display = a.Region
		} else {
			b.Display = a.Region + " (" + a.Comment + ")"
		}
	} else {
		if a.Comment == "" {
			b.Display = a.CountryCode + "-" + a.Region
		} else {
			b.Display = a.CountryCode + "-" + a.Region + " (" + a.Comment + ")"
		}
	}
	return b
}

func sortLoc(l *[]dexmodels.LocationYAML) {
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

func makeTitle(f dexmodels.EntryFields) string {

	var titlePart string
	if (f.Firstname + f.Lastname + f.Credentials) == "" {
		titlePart = f.Company
	} else {
		titlePart = strings.Join([]string{f.Firstname, f.Lastname, f.Credentials}, " ")
	}

	var title string
	if f.Country == "ZZ" {
		title = titlePart + fmt.Sprintf(" - %s from %s", f.Category, f.Region)
	} else {
		title = titlePart + fmt.Sprintf(" - %s from %s-%s", f.Category, f.Country, f.Region)
	}

	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, "  ", " ")
	return title
}

var regexInvalidPath = regexp.MustCompile("[^A-Za-z0-9_]+")

func makePath(f dexmodels.EntryFields) string {

	path := fmt.Sprintf("%d_%s-%s_%s",
		f.ID,
		strings.ToLower(f.Firstname),
		strings.ToLower(f.Lastname),
		strings.ToLower(f.Company),
	)

	// Remove diacritics from letters:
	// Cite: https://stackoverflow.com/questions/26722450/remove-diacritics-using-go
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	path, _, _ = transform.String(t, path)

	// Change runs of invalid chars to -
	path = regexInvalidPath.ReplaceAllString(path, "-")
	path = strings.TrimRight(path, "-_") // Clean up the end.

	return path
}

func getEntRaw(dbh reslist.Databaser) ([]dexmodels.Entry, error) {
	l, err := dbh.EntryList()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func tidyEnts(a []dexmodels.Entry, catMap map[int]dexmodels.CategoryYAML, locMap map[int]dexmodels.LocationYAML) []dexmodels.PathAndEntry {
	var b []dexmodels.PathAndEntry
	for _, j := range a {
		f := tidyEnt(j, catMap, locMap)
		f.Title = makeTitle(f)
		path := makePath(f)

		if j.Status != 1 || j.CategoryID == 0 {
			continue
		}
		b = append(b, dexmodels.PathAndEntry{
			Path:   path,
			Fields: f,
		})
	}
	return b
}

func tidyEnt(a dexmodels.Entry, catMap map[int]dexmodels.CategoryYAML, locMap map[int]dexmodels.LocationYAML) (b dexmodels.EntryFields) {
	b = dexmodels.EntryFields{}
	b.ID = a.ID
	b.Salutation = a.Salutation
	b.Firstname = a.Firstname
	b.Lastname = a.Lastname
	b.Credentials = a.Credentials
	b.JobTitle = a.JobTitle
	b.Company = a.Company
	b.ShortDesc = a.ShortDesc
	b.Phone = a.Phone
	b.Fax = a.Fax
	b.Address = a.Address
	b.Email = a.Email
	b.Email2 = a.Email2
	b.Website = a.Website
	b.Website2 = a.Website2
	b.Fees = a.Fees
	b.Description = a.Description

	b.Category = catMap[a.CategoryID].Name
	if a.LocationID == 0 {
		b.LocationDisplay = "Unknown"
		b.Country = "ZZ"
		b.Region = "Unknown"
	} else {
		//fmt.Printf("DEBUG: b.LocationDisplay %d\n", a.LocationID)
		b.LocationDisplay = locMap[a.LocationID].Display
		b.Country = locMap[a.LocationID].CountryCode
		b.Region = locMap[a.LocationID].Region
	}
	b.LastEditDate = a.LastEditDate

	return b
}

func sortEnt(l *[]dexmodels.PathAndEntry) {
	sort.Slice((*l), func(i, j int) bool {
		pi := (*l)[i].Fields.ID
		pj := (*l)[j].Fields.ID
		return pi < pj
	})
}
