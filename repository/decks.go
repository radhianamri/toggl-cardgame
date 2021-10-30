package repo

import (
	db "github.com/radhianamri/toggl-cardgame/lib/database"
	"github.com/radhianamri/toggl-cardgame/models/entities"
)

func GetDeckByID(deckID string) (entities.Deck, error) {
	var deck entities.Deck
	err := db.GetConn().Preload("Cards").Where("external_id = ?", deckID).Find(&deck).Error
	return deck, err
}
