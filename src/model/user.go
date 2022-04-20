package model

type User struct {
	Id    int `gorm:"primaryKey;autoIncrementIncrement"`
	Name  string
	Email string
	Phone *Phone
}
