package auth

import (
	"errors"
	"strings"
	"time"
)

// Username representa um nome de usuario valido
// Encapsula a string primitiva para garantir validacao e consistencia
type Username struct {
	value string
}

// ErrInvalidUsername é retornado quando um nome de usuario nao atende aos criterios
var ErrInvalidUsername = errors.New("nome de usuario invalido: deve ter entre 3 a 30 caracteres e conter apenas letras, numeros e underscores")

// NewUsername cria uma nova instancia de Username apos validacao
// Retorna erro se o nome de usuario nao valido
func NewUsername(username string) (Username, error) {
	trimmedUsername := strings.TrimSpace(username) // tira espaço

	if !isValidUsername(trimmedUsername) {
		return Username{}, ErrInvalidUsername
	}

	return Username{value: trimmedUsername}, nil
}

// String retorna a representação em string do nome de usuario
func (u Username) String() string {
	return u.value
}

// isValidUsername verifica se um nome de usuario atende aos criterio de validação
func isValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 30 {
		return false
	}

	for _, char := range username {
		if !isValidUsernameChar(char) {
			return false
		}
	}

	return true
}

// isValidUsernameChar verifica se um caractere é valido para nomes de usuario
func isValidUsernameChar(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		char == '_'
}

// PlayerID representa o indentificador unico de um jogador
type PlayerID struct {
	value uint64
}

// NewPlayerID cria uma nova instancia de PlayerID
func NewPlayerID(id uint64) PlayerID {
	return PlayerID{value: id}
}

// Value retorna o valor numerico do ID do jogador
func (id PlayerID) Value() uint64 {
	return id.value
}

// Player representa um jogador registrado no sistema
// Seguindo Object Calisthenics, limitamos a duas variaveis de instancia
// e encapsulamos os detalhes em tipos especificos
type Player struct {
	id        PlayerID
	username  Username
	email     Email
	password  PasswordHash
	createdAt time.Time
}

// NewPlayer cria uma nova instancia de Player
func NewPlayer(id PlayerID, username Username, email Email, password PasswordHash) Player {
	return Player{
		id:        id,
		username:  username,
		email:     email,
		password:  password,
		createdAt: time.Now(),
	}
}

// ID retorna o identificador unico do jogador
func (p Player) ID() PlayerID {
	return p.id
}

// Username retorna o nome de usuario do jogador.
func (p Player) Username() Username {
	return p.username
}

// Email retorna o email do jogador.
func (p Player) Email() Email {
	return p.email
}

// PasswordHash retorna o hash da senha do jogador
func (p Player) PasswordHash() PasswordHash {
	return p.password
}

// CreatedAt Retorna a data de criaçao da conta do jogador
func (p Player) CreatedAt() time.Time {
	return p.createdAt
}

// VerifyPassword verifica se uma senha em texto puro corresponde a senha do jogaodr
func (p Player) VerifiyPassword(plainPassword string) bool {
	return p.password.Compare(plainPassword)
}
