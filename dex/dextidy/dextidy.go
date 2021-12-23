package dextidy

//import "github.com/scrollodex/adminportal/dex/dexmodels"
import "github.com/scrollodex/adminportal/dex/dexmodels"

// MakeDisplayLoc returns the location as displayed to users.
func MakeDisplayLoc(loc dexmodels.Location) string {
	// This must match public/scrollodex.js:displayLocRenderer(). If you
	// change this, change it too.
	if loc.CountryCode == "ZZ" {
		if loc.Comment == "" {
			return loc.Region
		}
		return loc.Region + " (" + loc.Comment + ")"
	}
	if loc.Comment == "" {
		return loc.CountryCode + "-" + loc.Region
	}
	return loc.CountryCode + "-" + loc.Region + " (" + loc.Comment + ")"
}
