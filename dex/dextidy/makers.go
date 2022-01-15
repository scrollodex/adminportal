package dextidy

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/scrollodex/adminportal/dex/dexmodels"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func MakeTitle(f dexmodels.EntryFields) string {

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

func MakePathh(f dexmodels.EntryFields) string {

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
