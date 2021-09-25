package dexmodels

// MainListing represents the yaml output file delivered to Hugo.
// It is similar to forhugo.MainListing but uses the dexmodel structs.
type MainListing struct {
	Categories     []CategoryYAML `yaml:"categories"`
	Locations      []LocationYAML `yaml:"locations"`
	PathAndEntries []PathAndEntry `yaml:"entries"`
}

// PathAndEntry represents how an Entry is displayed in yaml.
type PathAndEntry struct {
	Path   string      `yaml:"path"` // {id}_{surname}-{prename}_{company-name}
	Fields EntryFields `yaml:"fields"`
}

// EntryFields represents the fields displayed in YAML.
type EntryFields struct {
	Title       string `yaml:"title"` // Generated {name or company} - {category} - {location}
	EntryCommon `yaml:",inline"`
	//
	Category        string `yaml:"categories"`
	LocationDisplay string `yaml:"location"`
	Country         string `yaml:"countries"`
	Region          string `yaml:"regions"`
	LastEditDate    string `yaml:"lastUpdate"`
}
