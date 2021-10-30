package decks

import "github.com/labstack/echo/v4"

func RegisterRoutes(r *echo.Group) {
	h := DeckHandler{}
	r.POST("/decks", h.CreateDeck)
	r.GET("/decks/:id", h.GetDeckByID)
	r.GET("/decks/:id/draw", h.DrawCards)
}
