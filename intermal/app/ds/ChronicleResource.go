package ds

import (
	"database/sql"
)

type ChronicleResource struct {
	ID                   uint           `gorm:"primaryKey;autoIncrement"`
	Image                string         `gorm:"type:varchar(255)"`
	Title                string         `gorm:"type:varchar(50);not null"`
	Author               string         `gorm:"type:varchar(50);not null"`
	DateOfCreation       string         `gorm:"type:varchar(50);not null"`
	TimeOfAction         string         `gorm:"type:varchar(50);not null"`
	Location             string         `gorm:"type:varchar(50);not null"`
	DetailedDescription  sql.NullString `gorm:"type:text;default:null"`
	DetailedSignificance sql.NullString `gorm:"type:text;default:null"`
	DetailedEditions     sql.NullString `gorm:"type:text;default:null"`
}
