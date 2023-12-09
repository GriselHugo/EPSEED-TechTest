package db

import (
	"time"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	ID        int 		`gorm:"primary_key"`
	Email     string 	`gorm:"unique;not null"`
	Username  string 	`gorm:"unique;not null"`
	Password  string 	`gorm:"not null"`
	Salt	  string 	`gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

const saltSize = 32

func GenerateSalt() (string, error) {
	saltBytes := make([]byte, saltSize)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(saltBytes), nil
}

func HashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))
	return hashedPassword
}

func CreateUser(username, password, email string) error {
	salt, err := GenerateSalt()
	if err != nil {
		return err
	}
	hashedPassword := HashPassword(password, salt)
	user := User{
		Username: username,
		Password: hashedPassword,
		Email: email,
		Salt: salt,
	}
	result := DbInstance.Create(&user)
	return result.Error
}

func GetUserByEmailandPassword(email, password string) (*User, error) {
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Email == email {
			if user.Password == HashPassword(password, user.Salt) {
				return &user, nil
			}
		}
	}
	return nil, nil
}

func GetUserByUsernameAndPassword(username, password string) (*User, error) {
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username {
			if user.Password == HashPassword(password, user.Salt) {
				return &user, nil
			}
		}
	}
	return nil, nil
}

func GetUserByUsername(username string) (*User, error) {
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, nil
}

func GetUserById(userID int) (*User, error) {
	var user User
	result := DbInstance.First(&user, userID)
	return &user, result.Error
}

func GetUsers() ([]User, error) {
	var users []User
	result := DbInstance.Find(&users)
	return users, result.Error
}

func UpdateUser(userID int, newUsername, neweEmail string) error {
	user, err := GetUserById(userID)
	if err != nil {
		return err
	}
	user.Username = newUsername
	user.Email = neweEmail
	result := DbInstance.Save(&user)
	return result.Error
}

func DeleteUser(userID int) error {
	result := DbInstance.Delete(&User{}, userID)
	return result.Error
}
