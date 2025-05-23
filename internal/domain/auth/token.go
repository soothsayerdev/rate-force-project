package auth

import (
	"errors"
	"time"
)

// Token representa um token JWT para autenticação
type Token struct {
	value     string
	expiresAt time.Time
}

// ErrInvalidToken é retornado quando um token é invalido ou expirou
var ErrInvalidToken = errors.New("token invalido/expirado")

// NewToken cria uma nova instancia de Token.
func NewToken(value string, expiresAt time.Time) Token {
	return Token{
		value:     value,
		expiresAt: expiresAt,
	}
}

// Value retorna o valor do token
func (t Token) Value() string {
	return t.value
}

// ExpiresAt retorna a data de expiração do token
func (t Token) ExpiresAt() time.Time {
	return t.expiresAt
}

// IsExpired verifica se o token expirou
func (t Token) isExpired() bool {
	return time.Now().After(t.expiresAt)
}
