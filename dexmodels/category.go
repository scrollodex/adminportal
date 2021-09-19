package dexmodels

// Category represents a single Category in storage.
type Category struct {
	ID          int    `yaml:"id"`
	Name        string `yaml:"category_name"`
	Description string `yaml:"description"`
	Priority    int    `yaml:"priority"`
	Icon        string `yaml:"icon"`
}

// CategoryYAML represents a single Category in YAML.
type CategoryYAML struct {
	Category `yaml:",inline"`
}
