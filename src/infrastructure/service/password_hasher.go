package service

import "golang.org/x/crypto/bcrypt"

type PasswordHasher struct {
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (ph *PasswordHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (ph *PasswordHasher) VerifyPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
