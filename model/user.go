package model

import "github.com/sunmi-OS/gocore/sqlite"

type User struct {
	Id       int64
	IsEnable int64
	Username string
	Password string
}

func (U *User) TableName() string {
	return "user"
}

func (U *User) CheckUsername(username string) (User, error) {

	user := User{}
	return user, sqlite.GetORM().Where("username = ?", username).First(&user).Error

}

func (U *User) CreateUser(username, password string) error {

	user := User{
		Username: username,
		Password: password,
		IsEnable: 1,
	}
	return sqlite.GetORM().Create(&user).Error

}
