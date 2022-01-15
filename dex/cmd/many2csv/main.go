package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	//"github.com/scrollodex/adminportal/dex/dexmodels"
	//"github.com/scrollodex/adminportal/dex/reslist"

	"github.com/scrollodex/adminportal/dex/reslist"
)

/*

What does it do?
It reads the DB and writes out the public.yaml file.

  makepublicyaml URL_TO_REPO filename

*/

func main() {
	// args:
	flag.Parse()
	repoURL := ""
	filenamePrefix := ""
	switch flag.NArg() {
	case 0:
		fmt.Println(flag.ErrHelp)
		os.Exit(1)
	case 1:
		repoURL = flag.Arg(0)
	case 2:
		repoURL = flag.Arg(0)
		filenamePrefix = flag.Arg(1)
	default:
		fmt.Println(flag.ErrHelp)
		os.Exit(1)
	}
	fmt.Printf("DEBUG: repoURL: %q\n", repoURL)

	dbh, err := reslist.New(repoURL, "")
	if err != nil {
		log.Fatal(err)
	}

	// GET THE DATA

	// Get cats
	rawCats, err := getCats(dbh)
	if err != nil {
		log.Fatal(err)
	}
	// Extract the strings
	stringCats := extractCats(rawCats)

	// Get locs
	rawLocs, err := getLocs(dbh)
	if err != nil {
		log.Fatal(err)
	} // Extract the strings
	stringLocs := extractLocs(rawLocs)

	// Index
	catMap := makeCatMap(rawCats)
	locMap := makeLocMap(rawLocs)

	// Get entries
	rawEntries, err := getEntries(dbh)
	if err != nil {
		log.Fatal(err)
	}
	stringEntries := extractEntries(rawEntries, catMap, locMap)

	// MODIFY DATA

	// WRITE OUT THE FILES

	writeCSV(filenamePrefix+"categories.csv", stringCats)
	writeCSV(filenamePrefix+"locations.csv", stringLocs)
	writeCSV(filenamePrefix+"entries.csv", stringEntries)

}
