package ds

type ChronicleResearch struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	IDRequestResearch uint   `gorm:"not null;uniqueIndex:idx_request_resource"`
	IDResource        uint   `gorm:"not null;uniqueIndex:idx_request_resource"`
	Quote             string `gorm:"type:text"`
	IsMatched         bool   `gorm:"type:boolean;default:false"`

	Request           RequestChronicleResearch `gorm:"foreignKey:IDRequestResearch"`
	ChronicleResource ChronicleResource        `gorm:"foreignKey:IDResource"`
}
