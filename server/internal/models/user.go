package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:current_timestamp(3)"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:NULL ON UPDATE current_timestamp(3)"`
	DeletedAt time.Time `json:"deletedAt" gorm:"default:NULL"`
}

func CreateUser(user *User) (*User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashed)
	res := db.Create(user)
	if res.RowsAffected == 0 {
		return &User{}, errors.New("can't create user")
	}
	return user, nil
}

func AuthUser(checkUser *User) (*User, error) {
	var user *User = &User{}
	res := db.First(&user, "username = ? ", checkUser.Username)
	if res.Error != nil && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(checkUser.Password)) != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
