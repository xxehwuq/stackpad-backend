package hash

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
