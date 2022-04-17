package model

type User struct {
	Id    int `gorm:"primaryKey;autoIncrement"`
	Name  string
	Email string
	Phone Phone
}
