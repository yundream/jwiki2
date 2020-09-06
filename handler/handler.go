package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"joinc.co.kr/jwiki/wiki"
)

// Handler ...
type Handler struct {
	wikiStore wiki.Store
}

// NewHandler ...
func NewHandler(ws wiki.Store) *Handler {
	return &Handler{
		wikiStore: ws,
	}
}

// Register ...
func (h *Handler) Register(api *echo.Group) {

	wiki := api.Group("/w/:filename")
	wiki.POST("", h.GetWikiFromID)
}

// StoreGet ...
func (h *Handler) StoreGet() {
	a, _ := h.wikiStore.GetByID(1)
	fmt.Println(a)
}

// GetWikiFromID ...
func (h *Handler) GetWikiFromID(c echo.Context) error {

	data := struct {
		Text string `json:"text"`
	}{
		Text: "Hello world",
	}
	return c.JSON(http.StatusCreated, data)
}
