package main

import (
	"fmt"
	"os"

	"github.com/scrollodex/adminportal/dex/reslist"
)

func main() {
	os.Setenv("ADMINPORTAL_BASEDIR", "/tmp")
	dbh, err := reslist.New(`git@scrollodex-github.com:scrollodex/beta-scollodex-db-%s.git`, "bi")
	fmt.Printf("ERROR=%v\n", err)
	fmt.Printf("  DBH=%v\n", dbh)
}
