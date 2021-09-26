package dexmodels

// Location represents a single Location in storage.
// json: Used to generate JSON for ZingGrid.
// yaml: Used to read/write n the reslist YAML database.
type Location struct {
	ID          int    `yaml:"id" json:"id"`
	CountryCode string `yaml:"country_code" json:"country_code"`
	Region      string `yaml:"region" json:"region"`
	Comment     string `yaml:"comment" json:"comment"`
}

// LocationYAML represents a single Location in YAML.
// yaml: Used as part of writing the Hugo YAML file (data/entries.yaml)
type LocationYAML struct {
	ID          int    `yaml:"id"`
	Display     string `yaml:"display"`
	CountryCode string `yaml:"country_code"`
	Region      string `yaml:"region"`
}
