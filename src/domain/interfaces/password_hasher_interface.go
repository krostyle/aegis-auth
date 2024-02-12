package interfaces

type PasswordHasherInterface interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, plainPassword string) bool
}
