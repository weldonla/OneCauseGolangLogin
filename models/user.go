package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
