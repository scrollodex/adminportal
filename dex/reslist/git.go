package reslist

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/scrollodex/adminportal/dex/dexmodels"
)

// GITHandle is the handle used to refer to GIT.
type GITHandle struct {
	url  string
	fdbh *FSHandle
}

// NewGIT creates a new GIT object.
func NewGit(url string) (Databaser, error) {
	db := &GITHandle{
		url: url,
	}

	// If directory exists, git pull. else git clone
	dir := whatDir(url)
	de, err := exists(dir)
	if err != nil {
		return nil, err
	}
	if de {
		fmt.Printf("DEBUG: DIR EXISTS: %v\n", dir)
		runCommand("git", "pull", "--force")
	} else {
		fmt.Printf("DEBUG: DIR DOES NOT EXIST: %v\n", dir)
		runCommand("git", "clone", url, dir)
	}

	// NewFS
	fdbh, err := NewFS(dir)
	db.fdbh = fdbh.(*FSHandle)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// whatDir reports the directory that "git clone" will create.
func whatDir(cs string) string {
	cs = strings.ReplaceAll(cs, ":", "_")
	cs = strings.ReplaceAll(cs, "@", "_")
	cs = strings.ReplaceAll(cs, "/", "_")
	cs = strings.ReplaceAll(cs, ".", "_")
	return cs
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func runCommand(name string, arg ...string) error {
	fmt.Printf("COMMAND: %s %v\n", name, arg)
	cmd := exec.Command(name, arg...)
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf(" OUTPUT: %s\n", stdoutStderr)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// CategoryStore stores a category in stable storage.
func (rh GITHandle) CategoryStore(data dexmodels.Category) error {

	// TODO: git pull
	// TODO: rh.fsthandle.CategoryStore(data)
	// TODO: git commit and push

	return nil
}

// LocationStore stores a location in stable storage.
func (rh GITHandle) LocationStore(data dexmodels.Location) error {

	// TODO: git pull
	// TODO: rh.fsthandle.LocationStore(data)
	// TODO: git commit and push

	return nil
}

// EntryStore stores an entry in stable storage.
func (rh GITHandle) EntryStore(data dexmodels.Entry) error {

	// TODO: git pull
	// TODO: rh.fsthandle.LocationStore(data)
	// TODO: git commit and push

	return nil
}

// CategoryList returns a list of all categories.
func (rh GITHandle) CategoryList() ([]dexmodels.Category, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.CategoryList(data)
	return nil, nil
}

// LocationList returns a list of all locations.
func (rh GITHandle) LocationList() ([]dexmodels.Location, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.LocationList(data)
	return nil, nil
}

// EntryList returns a list of all entries.
func (rh GITHandle) EntryList() ([]dexmodels.Entry, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.EntryList(data)
	return nil, nil
}

// CategoryGet gets a single item
func (rh GITHandle) CategoryGet(id int) (*dexmodels.Category, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.CategoryGet(data)
	return nil, nil
}

// LocationGet gets a single item
func (rh GITHandle) LocationGet(id int) (*dexmodels.Location, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.LocationGet(data)
	return nil, nil
}

// EntryGet gets a single item
func (rh GITHandle) EntryGet(id int) (*dexmodels.Entry, error) {
	// TODO: git pull
	// TODO: rh.fsthandle.EntryGet(data)
	return nil, nil
}
