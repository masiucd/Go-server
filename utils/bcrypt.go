package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword to hash users password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// ComparePassword to compare input password
func ComparePassword(password,hash string)bool  {
		err:= bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		return err != nil 
}