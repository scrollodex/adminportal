package reslist

//import "github.com/scrollodex/adminportal/dex/dexmodels"
import "github.com/scrollodex/adminportal/dex/dexmodels"

// Databaser is an interface for something that stores Scrollodex info.
type Databaser interface {
	CategoryStore(data dexmodels.Category) error
	CategoryGet(id int) (*dexmodels.Category, error)
	CategoryList() ([]dexmodels.Category, error)
	LocationStore(data dexmodels.Location) error
	LocationGet(id int) (*dexmodels.Location, error)
	LocationList() ([]dexmodels.Location, error)
	EntryStore(data dexmodels.Entry) error
	EntryGet(id int) (*dexmodels.Entry, error)
	EntryList() ([]dexmodels.Entry, error)
}
