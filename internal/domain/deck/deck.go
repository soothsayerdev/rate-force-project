package deck

import (
	"errors"
	"strings"
	"time"
)

// DeckID representa o identificador unico de um deck
type DeckID struct {
	value uint64
}

// NewDeckID cria uma nova instancia de DeckID
func NewDeckID(id uint64) DeckID {
	return DeckID{value: id}
}

// Value retorna o valor numerico do ID do deck
func (id DeckID) Value() uint64 {
	return id.value
}

// DeckName representa o nome de um deck
// Encapsula a string primitiva para garantir validação e consistencia
type DeckName struct {
	value string
}

// ErrInvalidDeckName é retornado quando um nome de deck nao atende aos criterios
var ErrInvalidDeckName = errors.New("nome do deck invalido: deve ter entre 2 e 150 caracteres")

// NewDeckName cria uma nova instancia de DeckName apos validação
// Retorna erro se o nome do deck nao for valido
func NewDeckName(name string) (DeckName, error) {
	trimmedName := strings.TrimSpace(name)

	if len(trimmedName) < 2 || len(trimmedName) > 150 {
		return DeckName{}, ErrInvalidDeckName
	}

	return DeckName{value: trimmedName}, nil
}

// String retorna a representação em string do nome do deck
func (n DeckName) String() string {
	return n.value
}

// DeckDescription representa a descrição de um deck
type DeckDescription struct {
	value string
}

// NewDeckDescription cria uma nova instancia de DeckDescription
func NewDeckDescription(description string) DeckDescription {
	return DeckDescription{value: strings.TrimSpace(description)}
}

// String retorna a representação em string da descrição do deck
func (d DeckDescription) String() string {
	return d.value
}

// Deck representa um deck no sistema
type Deck struct {
	id          DeckID
	name        DeckName
	description DeckDescription
	createdAt   time.Time
}

// NewDeck cria uma nova instancia de Deck
func NewDeck(id DeckID, name DeckName, description DeckDescription) Deck {
	return Deck{
		id:          id,
		name:        name,
		description: description,
		createdAt:   time.Now(),
	}
}

// ID retorna o identificador unico do deck
func (d Deck) ID() DeckID {
	return d.id
}

// Name retorna o nome do deck
func (d Deck) Name() DeckName {
	return d.name
}

// Description retorna a descrição do deck
func (d Deck) Description() DeckDescription {
	return d.description
}

// CreatedAt retorna a data de criação do deck
func (d Deck) CreatedAt() time.Time {
	return d.createdAt
}

// DeckList representa uma coleção de deck
type DeckList struct {
	items []Deck
}

// NewDeckList cria uma nova instancia de DeckList
func NewDeckList(decks []Deck) DeckList {
	return DeckList{items: decks}
}

// Items retorna todos os decks na lista
func (l DeckList) Items() []Deck {
	return l.items
}

// Count retorna o numero de decks na lista
func (l DeckList) Count() int {
	return len(l.items)
}

// FindByID encontra um deck pelo ID
func (l DeckList) FindByID(id DeckID) (Deck, bool) {
	for _, deck := range l.items {
		if deck.id.Value() == id.Value() {
			return deck, true
		}
	}

	return Deck{}, false
}
