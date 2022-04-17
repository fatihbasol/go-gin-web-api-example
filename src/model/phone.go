package model

type Phone struct {
	Id          int `gorm:"primaryKey;"`
	UserId      int `gorm:"unique"`
	CountryCode string
	Number      string
}
