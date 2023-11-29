package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
	Phone     string `json:"phone"`
}

func (user *User) SetPassword(password string) { // SetPassword is used to set the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password),14) // GenerateFromPassword is used to generate the hashed password
	// []byte(password) is used to convert the password to byte
	// 14 is the cost of the hashing algorithm
	user.Password = hashedPassword 

}