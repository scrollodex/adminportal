package reslist

import "github.com/scrollodex/scrollodex/dexmodels"

// Databaser is an interface for something that stores Scrollodex info.
type Databaser interface {
	CategoryStore(data dexmodels.Category) error
	CategoryList() ([]dexmodels.Category, error)
	LocationStore(data dexmodels.Location) error
	LocationList() ([]dexmodels.Location, error)
	EntryStore(data dexmodels.Entry) error
	EntryList() ([]dexmodels.Entry, error)
}
