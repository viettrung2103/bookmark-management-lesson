package service

import (
	"github.com/viettrung2103/bookmark-management-lesson/pkg/utils"
)

const (
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	passLength = 10
)

type genPassService struct {
	keygen utils.KeyGenerator
}

// GenPass represents the genpass service
//
//go:generate mockery --name=GenPass --filename=genpass.go
type GenPass interface {
	GeneratePassword() string
}

// NewGenPass return a GenPassService
func NewGenPass() GenPass {
	return &genPassService{}
}

// GeneratePassword generates a random password
func (s *genPassService) GeneratePassword() string {

	return utils.GenerateRandomString(passLength)

}
