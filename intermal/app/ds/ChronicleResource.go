package ds

type ChronicleResource struct {
	ID                   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Image                string `gorm:"type:varchar(255)" json:"image"`
	Title                string `gorm:"type:varchar(50);not null" json:"title"`
	Author               string `gorm:"type:varchar(50);not null" json:"author"`
	DateOfCreation       string `gorm:"type:varchar(50);not null" json:"date_of_creation"`
	TimeOfAction         string `gorm:"type:varchar(50);not null" json:"time_of_action"`
	Location             string `gorm:"type:varchar(50);not null" json:"location"`
	DetailedDescription  string `gorm:"type:text" json:"detailed_description"`
	DetailedSignificance string `gorm:"type:text" json:"detailed_significance"`
	DetailedEditions     string `gorm:"type:text" json:"detailed_editions"`
}
