package auth

// Auth contem entidade
// e regras de negocio relacionadas à autenticacao

import (
	"errors"
	"regexp"
	"strings"
)

// Email representa um endereco de email valido.
// Encapsula a string primitiva para garantir validacao e consistencia
type Email struct {
	value string
}

// ErrInvalidEmail é retornado quando um email não atende aos criterios de validacao
var ErrInvalidEmail = errors.New("email invalido")

// NewEmail cria uma nova instancia de Email apos validacao
// Retorna erro se o email nao for valido
func NewEmail(email string) (Email, error) {
	trimmedEmail := strings.TrimSpace(email)

	if !isValidEmail(trimmedEmail) {
		return Email{}, ErrInvalidEmail
	}

	return Email{value: trimmedEmail}, nil
}

// String retorna a representacao em string do email
func (e Email) String() string {
	return e.value
}

// isValidEmail verifica se um email é valido
func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	// expressão regular para validação de email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)

	return match
}
