package reslist

import (
	"fmt"
	"strings"
)

// New creates a new Databaser object based on urlpattern with
// substituting site where a %s appears.
func New(urlpattern string, site string) (Databaser, error) {

	fmt.Printf("RESLIST New(%q, %q)", urlpattern, site)

	cs := strings.ReplaceAll(urlpattern, "%s", site)

	if strings.HasPrefix(cs, "git@") {
		dbh, err := NewGit(cs)
		if err != nil {
			return nil, fmt.Errorf("reslist.NewGit(%q) failed: %w", cs, err)
		}
		return dbh, nil
	} else if strings.HasPrefix(cs, "file:") {
		cs = strings.TrimPrefix(cs, "file:")
		dbh, err := NewFS(cs)
		if err != nil {
			return nil, fmt.Errorf("reslist.NewFS(%q) failed: %w", cs, err)
		}
		return dbh, nil
	}
	dbh, err := NewFS(cs)
	if err != nil {
		return nil, fmt.Errorf("reslist.NewFS(%q) failed: %w", cs, err)
	}
	return dbh, nil
}
