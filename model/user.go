package model

import (
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

func (u *User) SaltAndSetPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = base64.StdEncoding.EncodeToString(hash)
}

func (u *User) CheckPassword(password string) (ok bool) {
	hashB, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		panic(err)
	}
	hashP := []byte(password)
	err = bcrypt.CompareHashAndPassword(hashB, hashP)
	if err != nil {
		return false
	}
	return true
}

func NewUserFromEmail(email string) *User {
	return &User{
		ID:    uuid.NewV4().String(),
		Email: strings.ToLower(email),
	}

}

func (u *User) GetUserMessage() *MessageUserV1 {
	result := MessageUserV1{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
	return &result

}
