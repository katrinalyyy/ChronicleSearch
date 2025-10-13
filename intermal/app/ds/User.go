package ds

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Email       string `gorm:"type:varchar(25);unique;not null" json:"login"`
	Name        string `gorm:"type:varchar(50);not null" json:"name"`
	Password    string `gorm:"type:varchar(50);not null" json:"password,omitempty"`
	IsModerator bool   `gorm:"type:boolean;default:false" json:"is_moderator,omitempty"`
}
