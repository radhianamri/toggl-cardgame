package decks

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/radhianamri/toggl-cardgame/lib/json"
	"github.com/radhianamri/toggl-cardgame/models/entities"
	"github.com/radhianamri/toggl-cardgame/models/views"
	repo "github.com/radhianamri/toggl-cardgame/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateDeck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/decks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := DeckHandler{}

	monkey.Patch(repo.GetCardsByCodesAndType, func([]string, string) ([]entities.Card, error) {
		return []entities.Card{
			{
				Value: "ACE",
				Suit:  "HEARTS",
				Code:  "AH",
			},
			{
				Value: "1",
				Suit:  "SPADES",
				Code:  "1S",
			},
		}, nil
	})
	monkey.Patch(repo.Save, func(interface{}) error {
		return nil
	})

	if assert.NoError(t, h.CreateDeck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp views.CreateDeckResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.Equal(t, 2, resp.Remaining)
		assert.Equal(t, false, resp.Shuffled)
	}
}

func TestCreateDeckShuffled(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("shuffled", "true")
	req := httptest.NewRequest(http.MethodPost, "/v1/decks?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := DeckHandler{}

	monkey.Patch(repo.GetCardsByCodesAndType, func([]string, string) ([]entities.Card, error) {
		return []entities.Card{
			{
				Value: "ACE",
				Suit:  "HEARTS",
				Code:  "AH",
			},
			{
				Value: "1",
				Suit:  "SPADES",
				Code:  "1S",
			},
		}, nil
	})
	monkey.Patch(repo.Save, func(interface{}) error {
		return nil
	})

	if assert.NoError(t, h.CreateDeck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp views.CreateDeckResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.Equal(t, 2, resp.Remaining)
		assert.Equal(t, true, resp.Shuffled)
	}
}

func TestCreateDeckInvalidCode(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("cards", "AH,1S")
	req := httptest.NewRequest(http.MethodPost, "/v1/decks?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := DeckHandler{}

	monkey.Patch(repo.GetCardsByCodesAndType, func([]string, string) ([]entities.Card, error) {
		return []entities.Card{
			{
				Value: "ACE",
				Suit:  "HEARTS",
				Code:  "AH",
			},
		}, nil
	})
	monkey.Patch(repo.Save, func(interface{}) error {
		return nil
	})

	if assert.NoError(t, h.CreateDeck(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetDeckByIDNotFound(t *testing.T) {
	e := echo.New()
	deckID := uuid.New().String()
	req := httptest.NewRequest(http.MethodGet, "/v1/decks/"+deckID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := DeckHandler{}

	monkey.Patch(repo.GetDeckByID, func(string) (entities.Deck, error) {
		return entities.Deck{}, gorm.ErrRecordNotFound
	})

	if assert.NoError(t, h.GetDeckByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestGetDeckByID(t *testing.T) {
	e := echo.New()
	deckID := uuid.New().String()
	req := httptest.NewRequest(http.MethodGet, "/v1/decks/"+deckID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := DeckHandler{}

	monkey.Patch(repo.GetDeckByID, func(string) (entities.Deck, error) {
		return entities.Deck{
			ExtID:    deckID,
			Shuffled: false,
			Cards: []entities.DeckCard{
				{
					CardID: 1,
					Base:   entities.Base{Status: true},
				},
				{
					CardID: 2,
					Base:   entities.Base{Status: true},
				},
			},
		}, nil
	})
	monkey.Patch(repo.GetCardsByIDList, func([]int) ([]entities.Card, error) {
		return []entities.Card{
			{
				Value: "ACE",
				Suit:  "SPADES",
				Code:  "AS",
			},
			{
				Value: "2",
				Suit:  "SPADES",
				Code:  "2S",
			},
		}, nil
	})

	if assert.NoError(t, h.GetDeckByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp views.DeckResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.Equal(t, 2, resp.Remaining)
		assert.Equal(t, false, resp.Shuffled)
		assert.Equal(t, "AS", resp.Cards[0].Code)
		assert.Equal(t, "2S", resp.Cards[1].Code)
	}
}

func TestDrawCardNotFound(t *testing.T) {
	e := echo.New()
	deckID := uuid.New().String()
	req := httptest.NewRequest(http.MethodGet, "/v1/decks/"+deckID+"draw", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := DeckHandler{}

	monkey.Patch(repo.GetDeckByID, func(string) (entities.Deck, error) {
		return entities.Deck{}, gorm.ErrRecordNotFound
	})

	if assert.NoError(t, h.DrawCards(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestDrawCardSuccess(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("count", "2")
	deckID := uuid.New().String()
	req := httptest.NewRequest(http.MethodGet, "/v1/decks/"+deckID+"draw?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := DeckHandler{}

	monkey.Patch(repo.GetDeckByID, func(string) (entities.Deck, error) {
		return entities.Deck{
			ExtID:    deckID,
			Shuffled: false,
			Cards: []entities.DeckCard{
				{
					CardID: 1,
					Base:   entities.Base{Status: true},
				},
				{
					CardID: 2,
					Base:   entities.Base{Status: true},
				},
			},
		}, nil
	})
	monkey.Patch(repo.GetNCardsFromDeck, func(entities.Deck, int) ([]entities.Card, error) {
		return []entities.Card{
			{
				Value: "ACE",
				Suit:  "SPADES",
				Code:  "AS",
			},
			{
				Value: "2",
				Suit:  "SPADES",
				Code:  "2S",
			},
		}, nil
	})

	if assert.NoError(t, h.DrawCards(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp []entities.Card
		_ = json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.Equal(t, 2, len(resp))
		assert.Equal(t, "AS", resp[0].Code)
		assert.Equal(t, "2S", resp[1].Code)
	}
}
