package model

type User struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Email    string
	Password string
}
