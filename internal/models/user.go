package models

import (
	"errors"
	"fmt"
	"time"

	"album-api/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Jwt                  string `json:"jwt"`

	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FindById(userId uint) (User, error) {
	var user User

	if err := db.First(&user, userId).Error; err != nil {
		return user, errors.New("User not found")
	}
	user.CleanFields()
	return user, nil
}

func (user *User) CleanFields() {
	user.Password = ""
	user.PasswordConfirmation = ""
}

func (user *User) SaveUser() (*User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.PasswordConfirmation = string(hashedPassword)

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		fmt.Printf("Cant generate JWT Token: %v", err)
	}

	user.Jwt = token

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Login(email string, password string) (string, error) {
	user := User{}
	err := db.Model(User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
