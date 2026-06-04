package service

import (
	"github.com/viettrung2103/bookmark-management-lesson/pkg/stringutils"
)

const (
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	passLength = 10
)

type passwordService struct{}

// GenPass represents the genpass service
//
//go:generate mockery --name=GenPass --filename=genpass.go
type Password interface {
	GeneratePassword() (string, error)
}

// NewGenPass return a GenPassService
func NewPassword() Password {
	return &passwordService{}
}

// GeneratePassword generates a random password
func (s *passwordService) GeneratePassword() (string, error) {

	return stringutils.GenerateCode(passLength)

}
