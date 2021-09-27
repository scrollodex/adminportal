package dextidy

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/scrollodex/scrollodex/reslist"
)

type NameVal = struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// CatNameVal returns the Categories as a list of Name/Value maps.
func CatNameVal(dbh reslist.Databaser) ([]NameVal, error) {
	orig, err := dbh.CategoryList()
	if err != nil {
		return nil, err
	}
	var nvl []NameVal
	for _, item := range orig {
		n := NameVal{Name: item.Name, Value: item.ID}
		nvl = append(nvl, n)
	}
	sort.Slice(nvl, func(i, j int) bool { return nvl[i].Name < nvl[j].Name })
	return nvl, nil
}

// GenCatList returns CatNameVal as a JSON string.
func GenCatList(dbh reslist.Databaser) (string, error) {
	nvl, err := CatNameVal(dbh)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(&nvl, "", "\t")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// LocNameVal returns the Locations as a list of Name/Value maps.
func LocNameVal(dbh reslist.Databaser) ([]NameVal, error) {
	orig, err := dbh.LocationList()
	if err != nil {
		return nil, err
	}
	var nvl []NameVal
	for _, item := range orig {
		n := NameVal{Name: MakeDisplayLoc(item), Value: item.ID}
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
	return nvl, nil
}

func GenLocList(dbh reslist.Databaser) (string, error) {
	nvl, err := LocNameVal(dbh)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(&nvl, "", "\t")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GenStatusList() (string, error) {
	nvl := []NameVal{
		{Name: "DISABLED", Value: 0},
		{Name: "active", Value: 1},
	}
	b, err := json.MarshalIndent(&nvl, "", "\t")
	if err != nil {
		return "", err
	}
	s := string(b)
	return s, nil
}
