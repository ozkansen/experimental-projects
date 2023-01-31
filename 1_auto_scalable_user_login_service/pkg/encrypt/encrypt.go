package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Encrypt interface {
	Encrypt(password []byte) ([]byte, error)
}
type Compare interface {
	Compare(hashedPassword []byte, password []byte) error
}

var _ Encrypt = (*Bcrypt)(nil)
var _ Compare = (*Bcrypt)(nil)

type Bcrypt struct {
	cost int
}

func NewBcrypt(cost int) *Bcrypt {
	if cost < bcrypt.MinCost {
		cost = bcrypt.MinCost
	} else if cost > bcrypt.MaxCost {
		cost = bcrypt.MaxCost
	}
	return &Bcrypt{cost: bcrypt.DefaultCost}
}

func (b *Bcrypt) Encrypt(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.cost)
}

func (b *Bcrypt) Compare(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
