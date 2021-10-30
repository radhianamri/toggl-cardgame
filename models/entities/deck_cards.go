package entities

type DeckCard struct {
	Base
	DeckID int64  `gorm:"column:deck_id" json:"-"`
	CardID int64  `gorm:"column:card_id" json:"-"`
	Cards  []Card `gorm:"foreignKey:ID;references:CardID"`
}

func (i DeckCard) TableName() string {
	return "deck_cards"
}
