package models

type User struct {
	ID           uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name         string  `gorm:"type:varchar(255)" json:"name"`
	Email        string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password     string  `gorm:"->;<-;not null" json:"-"`
	AccessToken  string  `gorm:"-" json:"accessToken,omitempty"`
	RefreshToken string  `gorm:"-" json:"refreshToken,omitempty"`
	Books        *[]Book `json:"books,omitempty"`
}
