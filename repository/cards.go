package repo

import (
	db "github.com/radhianamri/toggl-cardgame/lib/database"
	"github.com/radhianamri/toggl-cardgame/models/entities"
)

func GetCardsByIDList(idList []int) ([]entities.Card, error) {
	var cards []entities.Card
	err := db.GetConn().Where("id in (?)", idList).Find(&cards).Error
	return cards, err
}

func GetCardsByCodesAndType(codes []string, suitType string) ([]entities.Card, error) {
	var cards []entities.Card
	values := []interface{}{suitType}
	whereClause := "suit_type = ?"
	if len(codes) > 0 {
		whereClause += " AND code in (?)"
		values = append(values, codes)
	}
	err := db.GetConn().Where(whereClause, values...).Find(&cards).Error
	return cards, err
}

func GetNCardsFromDeck(deck entities.Deck, drawCount int) ([]entities.Card, error) {
	activeDeckCards := deck.GetRemaingCards()
	if drawCount > len(activeDeckCards) {
		drawCount = len(activeDeckCards)
	}
	deckCards := deck.Cards[0:drawCount]
	var cardIDList []int64
	for _, card := range deckCards {
		cardIDList = append(cardIDList, card.CardID)
	}
	var cards []entities.Card
	err := db.GetConn().Where("id in (?)", cardIDList).Find(&cards).Error
	if err != nil {
		return nil, err
	}

	return cards, RemoveCardsFromDeck(deck.ID, cardIDList)
}

func RemoveCardsFromDeck(deckID int64, cardIDList []int64) error {
	return db.GetConn().Model(&entities.DeckCard{}).
		Where("deck_id = ?", deckID).
		Where("card_id in (?)", cardIDList).
		Update("status", false).Error
}
