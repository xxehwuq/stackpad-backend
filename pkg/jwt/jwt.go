package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtManager interface {
	GenerateToken(id string) (string, error)
	GetClaimFromToken(token, key string) (interface{}, error)
}

type jwtManager struct {
	ttl time.Duration
	key string
}

func NewJwtManager(ttl time.Duration, key string) *jwtManager {
	return &jwtManager{
		ttl: ttl,
		key: key,
	}
}

func (m *jwtManager) GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(m.ttl).Unix(),
	})

	t, err := token.SignedString([]byte(m.key))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (m *jwtManager) GetClaimFromToken(token, key string) (interface{}, error) {
	claims := jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.key), nil
	})
	if !t.Valid {
		return nil, err
	}

	for claimsKey, value := range claims {
		if claimsKey == key {
			return value, nil
		}
	}

	return nil, nil
}
