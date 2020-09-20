package handler

import (
	"fmt"
	"net/http"

	"joinc.co.kr/jwiki/model"

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
	wiki.GET("", h.GetWikiFromID)
	wiki.POST("", h.CreateWiki)
}

// CreateWiki ...
func (h *Handler) CreateWiki(c echo.Context) error {
	fmt.Println("Create Wiki")
	var w model.Wiki
	req := &wikiCreateRequest{}
	if err := req.bind(c, &w); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	err := h.wikiStore.SaveWikiPage(&w)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return nil
}

// StoreGet ...
func (h *Handler) StoreGet() {
	a, _ := h.wikiStore.GetByID(1)
	fmt.Println(a)
}

// GetWikiFromID ...
func (h *Handler) GetWikiFromID(c echo.Context) error {
	data, _ := h.wikiStore.GetByID(1)
	return c.JSON(http.StatusOK, data)
}
