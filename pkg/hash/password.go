package hash

import (
	"crypto/sha256"
	"fmt"
)

type PasswordManager interface {
	Hash(password string) string
}

type passwordManager struct {
	salt string
}

func NewPasswordManager(salt string) *passwordManager {
	return &passwordManager{
		salt: salt,
	}
}

func (p *passwordManager) Hash(password string) string {
	hash := sha256.Sum256([]byte(password + p.salt))

	return fmt.Sprintf("%x", hash)
}
