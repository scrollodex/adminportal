package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/scrollodex/scrollodex/dexmodels"
	"github.com/scrollodex/scrollodex/reslist"
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
	outputFilename := ""
	switch flag.NArg() {
	case 0:
		fmt.Println(flag.ErrHelp)
		os.Exit(1)
	case 1:
		repoURL = flag.Arg(0)
	case 2:
		repoURL = flag.Arg(0)
		outputFilename = flag.Arg(1)
	default:
		fmt.Println(flag.ErrHelp)
		os.Exit(1)
	}
	fmt.Printf("DEBUG: repoURL: %q\n", repoURL)

	dbh, err := reslist.New(repoURL)
	if err != nil {
		log.Fatal(err)
	}

	rawCats, err := getCatRaw(dbh)
	if err != nil {
		log.Fatal(err)
	}
	cats := tidyCats(rawCats)
	sortCat(&cats)
	catMap := map[int]dexmodels.CategoryYAML{}
	for _, c := range cats {
		catMap[c.ID] = c
	}

	rawLocs, err := getLocRaw(dbh)
	if err != nil {
		log.Fatal(err)
	}
	locs := tidyLocs(rawLocs)
	sortLoc(&locs)
	locMap := map[int]dexmodels.LocationYAML{}
	for _, l := range locs {
		locMap[l.ID] = l
	}

	rawEnts, err := getEntRaw(dbh)
	if err != nil {
		log.Fatal(err)
	}
	ents := tidyEnts(rawEnts, catMap, locMap)
	sortEnt(&ents)

	// Make the yaml file for Hugo.
	yamlMaster := dexmodels.MainListing{
		Categories:     cats,
		Locations:      locs,
		PathAndEntries: ents,
	}
	d, err := yaml.Marshal(&yamlMaster)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	dStr := string(d)
	hugoYaml := "---\n" + dStr + "\n"
	if outputFilename == "" {
		fmt.Println(hugoYaml)
	} else {
		ioutil.WriteFile(outputFilename, []byte(hugoYaml), 0640)
		if err != nil {
			log.Fatalf("WriteFile %s: %v", outputFilename, err)
		}
	}

}
