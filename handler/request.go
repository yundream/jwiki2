package handler

import (
	"github.com/labstack/echo/v4"
	"joinc.co.kr/jwiki/model"
)

type wikiCreateRequest struct {
	Wiki struct {
		Name     string `json:"name" validate: "required"`
		Title    string `json:"title" validate: "required"`
		Author   string `json:"author" validate: "required"`
		Contents string `json:"contents" validate: "required"`
	} `json:"wiki"`
}

func (w *wikiCreateRequest) bind(c echo.Context, a *model.Wiki) error {
	if err := c.Bind(&w.Wiki); err != nil {
		return err
	}

	a.Name = w.Wiki.Name
	a.Title = w.Wiki.Title
	a.Author = w.Wiki.Author
	a.Contents = w.Wiki.Contents
	return nil
}
