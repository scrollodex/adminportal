package reslist

import (
	"fmt"
	"strings"
)

// New creates a new Databaser object based on URL.
func New(url string) (Databaser, error) {

	if strings.HasPrefix(url, "git@") {
		//dbh, err = NewGit(GitConfig{repoURL})
		return nil, fmt.Errorf("reslist.New(%q) failed: NOT IMPLEMENTED", url)
	} else if strings.HasPrefix(url, "/") || url != "" {
		dbh, err := NewFS(FSConfig{url})
		if err != nil {
			return nil, fmt.Errorf("reslist.New(%q) failed: %w", url, err)
		}
		return dbh, nil
	}

	return nil, fmt.Errorf("reslist.New(%q) failed: Invalid URL", url)
}
