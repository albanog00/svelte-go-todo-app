package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id" gorm:"primary_key index"`
	Username  string    `json:"username" gorm:"unique;index"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:current_timestamp(3)"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:NULL ON UPDATE current_timestamp(3)"`
	DeletedAt time.Time `json:"deletedAt" gorm:"default:NULL"`

	Task []Task `json:"tasks" gorm:"foreignKey:UserId"`
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

func AuthUser(user *User) error {
	res := db.First(&user, "username = ? ", user.Username)
	if res.Error != nil && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)) != nil {
		return errors.New("user not found")
	}
	return nil
}

func GetUser(userId string) (*User, error) {
	var user *User
	res := db.First(&user, "Id = ?", userId)
	if res.Error != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
