package dexmodels

// Entry represents a single Entry.
type Entry struct {
	EntryCommon `yaml:",inline"`
	//
	CategoryID          int    `yaml:"category_id"`
	LocationID          int    `yaml:"location_id"`
	Status              int    `yaml:"status"` // 0=Inactive, 1=Active
	LastEditDate        string `yaml:"lastUpdate"`
	PrivateLastEditBy   string `yaml:"private_last_edit_by"`
	PrivateAdminNotes   string `yaml:"private_admin_notes"`
	PrivateContactEmail string `yaml:"private_contact_email"`
}

// EntryCommon is the common fields among all Entry-like things.
type EntryCommon struct {
	ID          int    `yaml:"id"`
	Salutation  string `yaml:"salutation"`
	Firstname   string `yaml:"first_name"`
	Lastname    string `yaml:"last_name"`
	Credentials string `yaml:"credentials"`
	JobTitle    string `yaml:"job_title"`
	Company     string `yaml:"company"`
	ShortDesc   string `yaml:"short_desc"` // MarkDown (1 line)
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

// func (p *Entry) UnmarshalYAML(bytes []byte) error {
// 	fmt.Printf("ENTRY YES!!\n")
// 	log.Fatal("ENTRY YES")
// 	return nil
// }
// func (p *EntryCommon) UnmarshalYAML(bytes []byte) error {
// 	fmt.Printf("EntryCommon YES!!\n")
// 	log.Fatal("EntryCommon YES")
// 	return nil
// }
