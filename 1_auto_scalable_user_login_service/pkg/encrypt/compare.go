package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Compare interface {
	Compare(hashedPassword []byte, password []byte) error
}

var _ Compare = (*BcryptCompare)(nil)

type BcryptCompare struct{}

func NewBcryptCompare() *BcryptCompare {
	return &BcryptCompare{}
}

func (b *BcryptCompare) Compare(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
