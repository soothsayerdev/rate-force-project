package auth

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHash Representa uma senha hasheada
// Encapsula a string primitiva para garantir segurança
type PasswordHash struct {
	value string
}

// ErrWeakPassword é retornado quando uma senha nao atende aos criterios minimos
var ErrWeakPassword = errors.New("senha fraca: deve ter pelo menos 8 caracteres, incluindo letras maiusculas, minusculas e numeros")

// NewPasswordHash cria um novo hash de senha a partir de uma senha em texto puro.
// Retorna erro se a senha for fraca ou se ocorrer um erro no processo de hashing
func NewPasswordHash(plainPassword string) (PasswordHash, error) {
	if err := validatePassword(plainPassword); err != nil {
		return PasswordHash{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return PasswordHash{}, err
	}

	return PasswordHash{value: string(hash)}, nil
}

// String retorna a representação em string do hash da senha
func (p PasswordHash) String() string {
	return p.value
}

// Compare verifica se uma senha em texto puro corresponde ao hash armazenado
func (p PasswordHash) Compare(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(plainPassword))
	return err == nil
}

// validatePassword verifica se uma senha atende aos criterios minimos de segurança
func validatePassword(password string) error {
	if len(password) < 8 {
		return ErrWeakPassword
	}

	var hasUpper, hasLower, hasNumber bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber {
		return ErrWeakPassword
	}

	return nil
}
