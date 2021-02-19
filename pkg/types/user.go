package types

import (
	"golang.org/x/crypto/bcrypt"
)

// User type to handle auth
type User struct {
	ID       string  `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty" query:"id,omitempty"`
	Name     string  `json:"name" bson:"name" form:"name" query:"name"`
	Email    string  `json:"email" bson:"email" form:"email" query:"email" validate:"required"`
	Password string  `json:"password" bson:"password" form:"password" query:"password"`
	Causes   []Cause `json:"causes,omitempty" bson:"causes,omitempty"`
}

// RemovePassword Removes password from the struct
func (u *User) RemovePassword() {
	u.Password = ""
}

// EncryptPassword Encrypts the password
func (u *User) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// ComparePasswords compares two passwords
func (u *User) ComparePasswords(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	if err != nil {
		return false, err
	}
	return true, nil
}
