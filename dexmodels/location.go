package dexmodels

// Location represents a single Location in storage
type Location struct {
	ID          int    `yaml:"id"`
	CountryCode string `yaml:"country_code"`
	Region      string `yaml:"region"`
	Comment     string `yaml:"comment"`
}

// LocationYAML represents a single Location in YAML.
type LocationYAML struct {
	ID          int    `yaml:"id"`
	Display     string `yaml:"display"`
	CountryCode string `yaml:"country_code"`
	Region      string `yaml:"region"`
}
