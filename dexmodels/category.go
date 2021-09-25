package dexmodels

// Category represents a single Category in storage.
// json: Used when generating JSON for ZingGrid.
type Category struct {
	ID          int    `yaml:"id" json:"id"`
	Name        string `yaml:"category_name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Priority    int    `yaml:"priority" json:"-"`
	Icon        string `yaml:"icon" json:"icon"`
}

// CategoryYAML represents a single Category in YAML.
// yaml: Used when reading/writing raw data in the (yaml) database.
type CategoryYAML struct {
	Category `yaml:",inline"`
}
