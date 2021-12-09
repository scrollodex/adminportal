package reslist

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"

	//"github.com/scrollodex/adminportal/dex/dexmodels"
	"github.com/scrollodex/adminportal/dex/dexmodels"
)

// FSHandle is the handle used to refer to FS.
type FSHandle struct {
	directory string
}

// NewFS creates a new FS object.
func NewFS(dir string) (Databaser, error) {
	db := &FSHandle{
		directory: dir,
	}

	for _, n := range []string{"category", "location", "entry"} {
		err := os.MkdirAll(filepath.Join(db.directory, n), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db, nil
}

func (rh FSHandle) store(fieldName string, id int, y []byte) error {
	idStr := fmt.Sprintf("%05d", id)
	fn := filepath.Join(rh.directory, fieldName, idStr+".yaml")
	//fmt.Fprintf(os.Stderr, " %s", fn)
	fmt.Fprintf(os.Stderr, ".")
	return ioutil.WriteFile(fn, []byte("---\n"+string(y)), 0644)
}

// CategoryStore stores a category in stable storage.
func (rh FSHandle) CategoryStore(data dexmodels.Category) error {
	y, err := yaml.Marshal(&data)
	if err != nil {
		return err
	}
	return rh.store("category", data.ID, y)
}

// LocationStore stores a location in stable storage.
func (rh FSHandle) LocationStore(data dexmodels.Location) error {
	y, err := yaml.Marshal(&data)
	if err != nil {
		return err
	}
	return rh.store("location", data.ID, y)
}

// EntryStore stores an entry in stable storage.
func (rh FSHandle) EntryStore(data dexmodels.Entry) error {
	y, err := yaml.Marshal(&data)
	if err != nil {
		return err
	}
	return rh.store("entry", data.ID, y)
}

// CategoryList returns a list of all categories.
func (rh FSHandle) CategoryList() ([]dexmodels.Category, error) {
	fileSpec := filepath.Join(rh.directory, "category", "*.yaml")
	matches, err := filepath.Glob(fileSpec)
	sort.Strings(matches)
	if err != nil {
		return nil, err
	}
	var theList []dexmodels.Category
	fmt.Fprint(os.Stderr, "READ CATEGORIES: ")
	for _, match := range matches {
		//fmt.Fprintln(os.Stderr, match)
		fmt.Fprint(os.Stderr, ".")
		b, err := ioutil.ReadFile(match)
		if err != nil {
			return nil, err
		}
		var data dexmodels.Category
		yaml.Unmarshal(b, &data)
		chk := filepath.Join(rh.directory, "category",
			fmt.Sprintf("%05d.yaml", data.ID))
		if chk != match {
			log.Fatalf("File %s and the id: %d within does not match!", match, data.ID)
		}
		theList = append(theList, data)
	}
	fmt.Fprint(os.Stderr, "\n")
	return theList, nil
}

// LocationList returns a list of all locations.
func (rh FSHandle) LocationList() ([]dexmodels.Location, error) {
	fileSpec := filepath.Join(rh.directory, "location", "*.yaml")
	matches, err := filepath.Glob(fileSpec)
	sort.Strings(matches)
	if err != nil {
		return nil, err
	}
	fmt.Fprint(os.Stderr, "READ LOCATIONS: ")
	var theList []dexmodels.Location
	for _, match := range matches {
		//fmt.Fprintln(os.Stderr, match)
		fmt.Fprint(os.Stderr, ".")
		b, err := ioutil.ReadFile(match)
		if err != nil {
			return nil, err
		}
		var data dexmodels.Location
		yaml.Unmarshal(b, &data)
		chk := filepath.Join(rh.directory, "location",
			fmt.Sprintf("%05d.yaml", data.ID))
		if chk != match {
			log.Fatalf("File %s and the id: %d within does not match!", match, data.ID)
		}
		theList = append(theList, data)
	}
	fmt.Fprint(os.Stderr, "\n")
	return theList, nil
}

// EntryList returns a list of all entries.
func (rh FSHandle) EntryList() ([]dexmodels.Entry, error) {
	fileSpec := filepath.Join(rh.directory, "entry", "*.yaml")
	matches, err := filepath.Glob(fileSpec)
	sort.Strings(matches)
	if err != nil {
		return nil, err
	}
	var theList []dexmodels.Entry
	fmt.Fprint(os.Stderr, "READ ENTRIES: ")
	for _, match := range matches {
		//fmt.Fprint(os.Stderr, " ", match)
		fmt.Fprint(os.Stderr, ".")
		b, err := ioutil.ReadFile(match)
		if err != nil {
			return nil, err
		}
		var data dexmodels.Entry
		yaml.Unmarshal(b, &data)
		//fmt.Fprintf(os.Stderr, "DEBUG: data = %+v\n", data)
		chk := filepath.Join(rh.directory, "entry", fmt.Sprintf("%05d.yaml", data.ID))
		if chk != match {
			log.Fatalf("File %s and the id: %d within does not match!", match, data.ID)
		}
		theList = append(theList, data)
	}
	fmt.Fprint(os.Stderr, "\n")
	return theList, nil
}

func get(rh FSHandle, table string, id int, data interface{}) (interface{}, error) {
	fileSpec := filepath.Join(rh.directory, table, fmt.Sprintf("%05d.yaml", id))
	b, err := ioutil.ReadFile(fileSpec)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(b, data)
	// TODO(tlim): Verify data.ID == ID.
	//if data.ID != id {
	//	log.Fatalf("File %s and the id: %d within does not match!", fileSpec, id)
	//}
	return data, nil

}

// CategoryGet gets a single item
func (rh FSHandle) CategoryGet(id int) (*dexmodels.Category, error) {
	var data dexmodels.Category
	d, err := get(rh, "category", id, &data)
	return d.(*dexmodels.Category), err
}

// LocationGet gets a single item
func (rh FSHandle) LocationGet(id int) (*dexmodels.Location, error) {
	var data dexmodels.Location
	d, err := get(rh, "location", id, &data)
	return d.(*dexmodels.Location), err
}

// EntryGet gets a single item
func (rh FSHandle) EntryGet(id int) (*dexmodels.Entry, error) {
	var data dexmodels.Entry
	d, err := get(rh, "entry", id, &data)
	fmt.Fprintf(os.Stderr, "DEBUG FS ENTRY = %v\n", d)
	return d.(*dexmodels.Entry), err
}
