package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"joinc.co.kr/jwiki/model"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
	"joinc.co.kr/jwiki/router"
	"joinc.co.kr/jwiki/store"
	"joinc.co.kr/jwiki/wiki"
)

var (
	ws wiki.Store
	h  *Handler
	e  *echo.Echo
)

func setup() {
	ws = store.NewWikiStore("joinc")
	h = NewHandler(ws)
	e = router.New()
}

func Test_GetWikiFromID(t *testing.T) {
	setup()
	req := httptest.NewRequest(echo.POST, "/api/w/hello", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.GetWikiFromID(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var w model.Wiki
		err := json.Unmarshal(rec.Body.Bytes(), &w)
		assert.NoError(t, err)
		assert.Equal(t, w.Name, "yundream")
	}
}
