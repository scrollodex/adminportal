package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mehanizm/airtable" // https://pkg.go.dev/github.com/mehanizm/airtable#section-readme
	"github.com/scrollodex/adminportal/dex/dexmodels"
	"gopkg.in/yaml.v3"
)

func main() {
	// args:
	flag.Parse()
	outputFilename := "entries.yaml"
	switch flag.NArg() {
	case 0:
		//fmt.Println(flag.ErrHelp)
		//os.Exit(1)
	case 1:
		outputFilename = flag.Arg(0)
	default:
		fmt.Println(flag.ErrHelp)
		os.Exit(1)
	}

	// Gather data:
	client := airtable.NewClient(os.Getenv("AIRTABLE_APIKEY"))
	locTable := client.GetTable(os.Getenv("AIRTABLE_BASE_ID"), "Locations")
	catTable := client.GetTable(os.Getenv("AIRTABLE_BASE_ID"), "Categories")
	entTable := client.GetTable(os.Getenv("AIRTABLE_BASE_ID"), "Entries")
	locAir := getRecordsAll(locTable)
	catAir := getRecordsAll(catTable)
	entAir := getRecordsAll(entTable)
	fmt.Printf("LOC-AIR %d\n", len(locAir))
	fmt.Printf("CAT-AIR %d\n", len(catAir))
	fmt.Printf("ENT-AIR %d\n", len(entAir))
	fmt.Println()

	//var cats []dexmodels.CategoryYAML
	cats := extractCats(catAir)
	locs := extractLocs(locAir)
	ents := extractEnts(entAir)

	sortCats(&cats)
	sortLocs(&locs)
	sortEnts(&ents)

	fmt.Printf("LOC-YAML %d\n", len(locs))
	fmt.Printf("CAT-YAML %d\n", len(cats))
	fmt.Printf("ENT-YAML %d\n", len(ents))
	fmt.Println()

	// Generate yaml
	hugoYaml := getYaml(cats, locs, ents)
	err := ioutil.WriteFile(outputFilename, []byte(hugoYaml), 0666)
	if err != nil {
		log.Fatalf("WriteFile %s: %v", outputFilename, err)
	}
}

// getYaml generates the full YAML file that Hugo expects.
func getYaml(
	cats []dexmodels.CategoryYAML,
	locs []dexmodels.LocationYAML,
	ents []dexmodels.PathAndEntry,
) string {

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
	return "---\n" + dStr + "\n"
}

// getRecordsAll gets all records of a table (all pages).
func getRecordsAll(table *airtable.Table) []*airtable.Record {
	// TODO(tlim): If we are rate limited, retry.

	var result []*airtable.Record

	var offset string
	for {
		// Get 1 page of records.
		records, err := table.GetRecords().
			WithOffset(offset).
			Do()
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, records.Records...)

		// Stop when we're out of records.
		offset = records.Offset
		if offset == "" {
			break
		}
	}

	return result
}
