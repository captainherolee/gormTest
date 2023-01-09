package models

type User struct {
	ID           int    `json:"userId" gorm:"column:user_id;AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Email        string `json:"email" gorm:"column:user_email;size:100;unique;not null"`
	Name         string `json:"name" gorm:"column:user_name;size:100"`
	Organization string `json:"organization" gorm:"column:organization;size:100"`
	Tag          string `json:"tag" gorm:"column:tag;size:100"`
	CreatedAt    int    `json:"createdAt" gorm:"column:created_at;type:int(13)"`
	UpdatedAt    int    `json:"updatedAt" gorm:"column:updated_at;type:int(13)"`
}

// TableName gets table name
func (u *User) TableName() string {
	return "user"
}
