package dexmodels

// Category represents a single Category in storage.
// json: Used to generate JSON for ZingGrid.
// yaml: Used to read/write n the reslist YAML database.
type Category struct {
	ID          int    `yaml:"id" json:"id"`
	Name        string `yaml:"category_name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Priority    int    `yaml:"priority" json:"-"`
	Icon        string `yaml:"icon" json:"icon"`
}

// CategoryYAML represents a single Category in YAML.
// yaml: Used as part of writing the Hugo YAML file (data/entries.yaml)
type CategoryYAML struct {
	Category `yaml:",inline"`
}
