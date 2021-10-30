package views

import "github.com/radhianamri/toggl-cardgame/models/entities"

type DeckResponse struct {
	DeckID    string          `json:"deck_id"`
	Shuffled  bool            `json:"shuffled"`
	Remaining int             `json:"remaining"`
	Cards     []entities.Card `json:"cards"`
}
