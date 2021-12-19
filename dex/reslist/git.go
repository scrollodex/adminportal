package reslist

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/scrollodex/adminportal/dex/dexmodels"
)

// GITHandle is the handle used to refer to GIT.
type GITHandle struct {
	url      string
	dir      string
	fshandle *FSHandle
}

// Initialization

// NewGIT creates a new GIT object.
func NewGit(url string) (Databaser, error) {
	db := &GITHandle{
		url: url,
		dir: cloneDirName(url),
	}

	db.init()

	// NewFS
	fshandle, err := NewFS(db.dir)
	db.fshandle = fshandle.(*FSHandle)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (rh GITHandle) init() error {
	url := rh.url
	dir := rh.dir
	basedir := filepath.Join(os.Getenv("ADMINPORTAL_BASEDIR"))
	repodir := filepath.Join(basedir, dir)

	// Are we already cloned and ready?  Just "git pull".
	err := os.Chdir(repodir)
	if err == nil {
		fmt.Printf("DEBUG: REPO DIR EXISTS. PULLING: %v\n", dir)
		return runCommand("git", "pull", "--force")
	}

	// Otherwise, we have to "git clone" and cd into it.
	err = os.Chdir(basedir)
	if err != nil {
		return fmt.Errorf("chdir(%q) failed: %w", basedir, err)
	}

	fmt.Printf("DEBUG: DIR DOES NOT EXIST: %v\n", dir)
	if err = runCommand("git", "clone", url, dir); err != nil {
		return err
	}

	return os.Chdir(repodir)
}

func (rh GITHandle) commit() error {
	fmt.Printf("DEBUG: COMMITTING\n")
	return runCommand("git", "commit", "-m", "Automated commit from AdminPanel")
}

// cloneDirName reports the directory that "git clone" will create.
func cloneDirName(cs string) string {
	cs = strings.ReplaceAll(cs, ":", "_")
	cs = strings.ReplaceAll(cs, "@", "_")
	cs = strings.ReplaceAll(cs, "/", "_")
	cs = strings.ReplaceAll(cs, ".", "_")
	return cs
}

//// dirExists returns whether the given file or directory exists
//func dirExists(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}

func runCommand(name string, arg ...string) error {
	fmt.Printf("COMMAND: %s %v\n", name, arg)
	cmd := exec.Command(name, arg...)
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf(" OUTPUT: %s\n", stdoutStderr)
	return err
}

// Store

// CategoryStore stores a category in stable storage.
func (rh GITHandle) CategoryStore(data dexmodels.Category) error {
	if err := rh.init(); err != nil {
		return err
	}
	if err := rh.fshandle.CategoryStore(data); err != nil {
		return err
	}
	return rh.commit()
}

// LocationStore stores a location in stable storage.
func (rh GITHandle) LocationStore(data dexmodels.Location) error {
	if err := rh.init(); err != nil {
		return err
	}
	if err := rh.fshandle.LocationStore(data); err != nil {
		return err
	}
	return rh.commit()
}

// EntryStore stores an entry in stable storage.
func (rh GITHandle) EntryStore(data dexmodels.Entry) error {
	if err := rh.init(); err != nil {
		return err
	}
	if err := rh.fshandle.EntryStore(data); err != nil {
		return err
	}
	return rh.commit()
}

// List

// CategoryList returns a list of all categories.
func (rh GITHandle) CategoryList() ([]dexmodels.Category, error) {
	if err := rh.init(); err != nil {
		return nil, err
	}
	return rh.fshandle.CategoryList()
}

// LocationList returns a list of all locations.
func (rh GITHandle) LocationList() ([]dexmodels.Location, error) {
	if err := rh.init(); err != nil {
		return nil, err
	}
	return rh.fshandle.LocationList()
}

// EntryList returns a list of all entries.
func (rh GITHandle) EntryList() ([]dexmodels.Entry, error) {
	if err := rh.init(); err != nil {
		return nil, err
	}
	return rh.fshandle.EntryList()
}

// Get

// CategoryGet gets a single item
func (rh GITHandle) CategoryGet(id int) (*dexmodels.Category, error) {
	if err := rh.init(); err != nil {
		return nil, err
	}
	return rh.fshandle.CategoryGet(id)
}

// LocationGet gets a single item
func (rh GITHandle) LocationGet(id int) (*dexmodels.Location, error) {
	if err := rh.init(); err != nil {
		return nil, err
	}
	return rh.fshandle.LocationGet(id)
}

// EntryGet gets a single item
func (rh GITHandle) EntryGet(id int) (*dexmodels.Entry, error) {
	if err := rh.init(); err != nil {
		return nil, err
	}
	return rh.fshandle.EntryGet(id)
}
