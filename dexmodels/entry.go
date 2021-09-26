package dexmodels

// Entry represents a single Entry.
// yaml: Used when reading/writing raw data in the (yaml) database.
// json: used when generating JSON for ZingGrid.
type Entry struct {
	EntryCommon `yaml:",inline"`
	//
	CategoryID          int    `yaml:"category_id" json:"category_id"`
	LocationID          int    `yaml:"location_id" json:"location_id"`
	Status              int    `yaml:"status"` // 0=Inactive, 1=Active
	LastEditDate        string `yaml:"lastUpdate" json:"last_update"`
	PrivateLastEditBy   string `yaml:"private_last_edit_by" json:"private_last_edit_by"`
	PrivateAdminNotes   string `yaml:"private_admin_notes" json:"private_admin_notes"`
	PrivateContactEmail string `yaml:"private_contact_email" json:"private_contact_email"`
}

// EntryCommon is the common fields among all Entry-like things.
type EntryCommon struct {
	ID          int    `yaml:"id"`
	Salutation  string `yaml:"salutation"`
	Firstname   string `yaml:"first_name" json:"first_name"`
	Lastname    string `yaml:"last_name" json:"last_name"`
	Credentials string `yaml:"credentials"`
	JobTitle    string `yaml:"job_title" json:"job_title"`
	Company     string `yaml:"company" json:"company"`
	ShortDesc   string `yaml:"short_desc" json:"short_desc"` // MarkDown (1 line)
	Phone       string `yaml:"phone"`
	Fax         string `yaml:"fax"`
	Address     string `yaml:"address"`
	Email       string `yaml:"email"`
	Email2      string `yaml:"email2"`
	Website     string `yaml:"website"`
	Website2    string `yaml:"website2"`
	Fees        string `yaml:"fees"`        // MarkDown
	Description string `yaml:"description"` // MarkDown
}
