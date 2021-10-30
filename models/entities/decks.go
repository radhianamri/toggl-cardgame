package entities

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deck struct {
	Base
	ExtID    string     `gorm:"column:external_id" json:"deck_id"`
	Shuffled bool       `gorm:"column:shuffled" json:"shuffled"`
	Cards    []DeckCard `gorm:"foreignKey:DeckID" json:"-"`
}

func (i Deck) TableName() string {
	return "decks"
}

func (i *Deck) BeforeSave(tx *gorm.DB) error {
	i.ExtID = uuid.New().String()
	i.PrePersist()
	for x := range i.Cards {
		i.Cards[x].PrePersist()
	}
	return nil
}

func (i *Deck) AddCards(cards []Card) {
	for _, card := range cards {
		i.Cards = append(i.Cards, DeckCard{
			CardID: card.ID,
			Base: Base{
				Status: true,
			},
		})
	}
}

func (i *Deck) Shuffle() {
	a := i.Cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	i.Shuffled = true
}

func (i *Deck) GetRemaingCards() (activeDeckCards []DeckCard) {
	for _, card := range i.Cards {
		if card.Status {
			activeDeckCards = append(activeDeckCards, card)
		}
	}
	return
}
