package service

type TokenGeneratorInterface interface {
	GenerateToken(userID string) (string, error)
}
