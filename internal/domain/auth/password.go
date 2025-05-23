package auth

// PasswordHash Representa uma senha hasheada
// Encapsula a string primitiva para garantir segurança
type PasswordHash struct {
	value string
}