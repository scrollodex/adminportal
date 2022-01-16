package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/scrollodex/adminportal/dex/dexmodels"
	"github.com/scrollodex/adminportal/dex/dextidy"
	"github.com/scrollodex/adminportal/dex/reslist"
)

// Categories

func getCats(dbh reslist.Databaser) ([]dexmodels.Category, error) {
	l, err := dbh.CategoryList()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func extractCats(items []dexmodels.Category) [][]string {
	var result [][]string
	result = append(result, []string{
		"Name",
		"IconFilename",
		"Description",
		"x-CategoryID",
		"x-Priority",
	})

	for _, r := range items {
		row := []string{
			r.Name,                        // Name        string `yaml:"category_name" json:"name"`
			r.Icon,                        // Icon        string `yaml:"icon" json:"icon"`
			r.Description,                 // Description string `yaml:"description" json:"description"`
			fmt.Sprintf("%d", r.ID),       // ID          int    `yaml:"id" json:"id"`
			fmt.Sprintf("%d", r.Priority), // Priority    int    `yaml:"priority" json:"-"`
		}
		result = append(result, row)
	}
	return result
}

// Locations

func getLocs(dbh reslist.Databaser) ([]dexmodels.Location, error) {
	l, err := dbh.LocationList()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func extractLocs(items []dexmodels.Location) [][]string {
	var result [][]string
	result = append(result, []string{
		"Location",
		"x-LocationID",
		"x-CountryCode",
		"x-Region",
		"x-Comment",
	})
	for _, r := range items {
		row := []string{
			dextidy.MakeDisplayLoc(r),
			fmt.Sprintf("%d", r.ID), // ID          int    `yaml:"id" json:"id"`
			r.CountryCode,           // CountryCode string `yaml:"country_code" json:"country_code"`
			r.Region,                // Region      string `yaml:"region" json:"region"`
			r.Comment,               // Comment     string `yaml:"comment" json:"comment"`
		}
		result = append(result, row)
	}
	return result
}

// Entries

func getEntries(dbh reslist.Databaser) ([]dexmodels.Entry, error) {
	l, err := dbh.EntryList()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func extractEntries(
	items []dexmodels.Entry,
	catMap map[int]string,
	locMap map[int]string,
) [][]string {
	var result [][]string
	result = append(result, []string{
		"EntryID",
		"Status",
		"Category",
		"Location",
		"Sal",
		"First",
		"Last",
		"Suffix",
		"Job_Title",
		"Company",
		"Short Description",
		"Phone",
		"Fax",
		"Address",
		"Email",
		"Email2",
		"Website",
		"Website2",
		"Description",
		"Fees",
		//
		"PRIVATE_admin_notes",
		"PRIVATE_contact_email",
		//
		"x-lastUpdate",
		"x-private_last_edit_by",
	})

	for _, r := range items {

		// Status              int    `yaml:"status"` // 0=Inactive, 1=Active
		status := "SHOW"
		if r.Status != 1 {
			status = "HIDDEN"
		}

		// CategoryID          int    `yaml:"category_id" json:"category_id"`
		//catid := fmt.Sprintf("%d", r.CategoryID)
		catid, ok := catMap[r.CategoryID]
		if !ok {
			fmt.Printf("No such category_id=%d\n", r.CategoryID)
			catid = fmt.Sprintf("%d", r.CategoryID)
		}

		// LocationID          int    `yaml:"location_id" json:"location_id"`
		//locid := fmt.Sprintf("%d", r.LocationID)
		locid, ok := locMap[r.LocationID]
		if !ok {
			fmt.Printf("No such location_id=%d\n", r.LocationID)
			locid = fmt.Sprintf("%d", r.LocationID)
		}

		row := []string{
			fmt.Sprintf("%d", r.ID), // ID          int    `yaml:"id"`
			// Foreign Fields
			status,
			catid,
			locid,

			r.Salutation,  // Salutation  string `yaml:"salutation"`
			r.Firstname,   // Firstname   string `yaml:"first_name" json:"first_name"`
			r.Lastname,    // Lastname    string `yaml:"last_name" json:"last_name"`
			r.Credentials, // Credentials string `yaml:"credentials"`
			r.JobTitle,    // JobTitle    string `yaml:"job_title" json:"job_title"`
			r.Company,     // Company     string `yaml:"company" json:"company"`
			r.ShortDesc,   // ShortDesc   string `yaml:"short_desc" json:"short_desc"` // MarkDown (1 line)
			r.Phone,       // Phone       string `yaml:"phone"`
			r.Fax,         // Fax         string `yaml:"fax"`
			r.Address,     // Address     string `yaml:"address"`
			r.Email,       // Email       string `yaml:"email"`
			r.Email2,      // Email2      string `yaml:"email2"`
			r.Website,     // Website     string `yaml:"website"`
			r.Website2,    // Website2    string `yaml:"website2"`
			r.Description, // Description string `yaml:"description"` // MarkDown
			r.Fees,        // Fees        string `yaml:"fees"`        // MarkDown

			// More
			r.PrivateAdminNotes,   // PrivateAdminNotes   string `yaml:"private_admin_notes" json:"private_admin_notes"`
			r.PrivateContactEmail, // PrivateContactEmail string `yaml:"private_contact_email" json:"private_contact_email"`
			//
			r.LastEditDate,      // LastEditDate        string `yaml:"lastUpdate" json:"last_update"`
			r.PrivateLastEditBy, // PrivateLastEditBy   string `yaml:"private_last_edit_by" json:"private_last_edit_by"`

		}
		result = append(result, row)
	}
	return result
}

func writeCSV(filename string, data [][]string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Write out CSV:
	w := csv.NewWriter(f)
	w.WriteAll(data)
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func makeCatMap(items []dexmodels.Category) map[int]string {
	r := map[int]string{}
	for _, item := range items {
		r[item.ID] = item.Name
	}
	return r
}

func makeLocMap(items []dexmodels.Location) map[int]string {
	r := map[int]string{}
	for _, item := range items {
		//fmt.Printf("r[%d] = %q\n", item.ID, dextidy.MakeDisplayLoc(item))
		r[item.ID] = dextidy.MakeDisplayLoc(item)
	}
	return r
}
