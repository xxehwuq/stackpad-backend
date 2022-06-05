package entity

import (
	"github.com/yaroslavyarosh/stackpad-backend/pkg/hash"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/jwt"
)

type Pkg struct {
	PasswordManager hash.PasswordManager
	JwtManager      jwt.JwtManager
}
