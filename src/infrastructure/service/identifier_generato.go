package service

import "github.com/google/uuid"

type UUIDGenerator struct{}

func (u *UUIDGenerator) GenerateUUID() string {
	return uuid.New().String()
}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}
