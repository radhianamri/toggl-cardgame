package decks

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/radhianamri/toggl-cardgame/enums"
	"github.com/radhianamri/toggl-cardgame/lib/log"
	"github.com/radhianamri/toggl-cardgame/lib/response"
	"github.com/radhianamri/toggl-cardgame/models/entities"
	"github.com/radhianamri/toggl-cardgame/models/views"
	repo "github.com/radhianamri/toggl-cardgame/repository"
)

type DeckHandler struct {
}

func (d *DeckHandler) CreateDeck(c echo.Context) error {
	suitType := enums.FrechSuit
	if enums.CardSuitMap[c.QueryParam("suit_type")] {
		suitType = c.QueryParam("suit_type")
	}
	var codes []string
	if c.QueryParam("cards") != "" {
		codes = strings.Split(c.QueryParam("cards"), ",")
	}
	cards, err := repo.GetCardsByCodesAndType(codes, suitType)
	if err != nil {
		log.Errorf("Error fetching cards by suit %s", err.Error())
		return response.InternalServerError(c)
	}
	if len(codes) != 0 && len(codes) != len(cards) {
		return response.BadRequest(c, "invalid card codes")
	}

	var newDeck entities.Deck
	newDeck.AddCards(cards)
	shuffled, _ := strconv.ParseBool(c.QueryParam("shuffled"))
	if shuffled {
		newDeck.Shuffle()
	}

	if err := repo.Save(&newDeck); err != nil {
		log.Errorf("Error saving deck %s", err.Error())
		return response.InternalServerError(c)
	}

	resp := views.CreateDeckResponse{
		DeckID:    newDeck.ExtID,
		Shuffled:  newDeck.Shuffled,
		Remaining: len(cards),
	}
	return c.JSON(http.StatusOK, resp)
}

func (d *DeckHandler) GetDeckByID(c echo.Context) error {
	deckID := c.Param("id")
	deck, err := repo.GetDeckByID(deckID)
	if err != nil {
		return response.NotFound(c, "Invalid Card ID")
	}

	var cardIDList []int
	for _, card := range deck.GetRemaingCards() {
		cardIDList = append(cardIDList, int(card.CardID))
	}

	cards, err := repo.GetCardsByIDList(cardIDList)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	deckResponse := views.DeckResponse{
		DeckID:    deck.ExtID,
		Shuffled:  deck.Shuffled,
		Remaining: len(cards),
		Cards:     cards,
	}
	return c.JSON(http.StatusOK, deckResponse)
}

func (d *DeckHandler) DrawCards(c echo.Context) error {
	deckID := c.Param("id")
	deck, err := repo.GetDeckByID(deckID)
	if err != nil {
		return response.ErrOrNotFound(c, err, "Invalid Card ID")
	}

	drawCountParam := c.QueryParam("count")
	drawCount, err := strconv.Atoi(drawCountParam)
	if err != nil {
		return response.BadRequest(c, "Invalid amount of cards to draw")
	}

	cards, err := repo.GetNCardsFromDeck(deck, drawCount)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cards)
}
