package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Encrypt interface {
	Encrypt(password []byte) ([]byte, error)
}

var _ Encrypt = (*BcryptEncrypt)(nil)

type BcryptEncrypt struct {
	cost int
}

func NewBcryptEncrypt(cost int) *BcryptEncrypt {
	if cost < bcrypt.MinCost {
		cost = bcrypt.MinCost
	} else if cost > bcrypt.MaxCost {
		cost = bcrypt.MaxCost
	}
	return &BcryptEncrypt{cost: bcrypt.DefaultCost}
}

func (b *BcryptEncrypt) Encrypt(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.cost)
}
