package entity

import (
	"github.com/yaroslavyarosh/stackpad-backend/pkg/hash"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/jwt"
	"github.com/yaroslavyarosh/stackpad-backend/pkg/mail"
)

type Pkg struct {
	PasswordManager hash.PasswordManager
	JwtManager      jwt.JwtManager
	MailManager     mail.MailManager
}
