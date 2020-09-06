package main

import (
	"fmt"

	"joinc.co.kr/jwiki/handler"
	"joinc.co.kr/jwiki/router"
	"joinc.co.kr/jwiki/store"
)

func main() {
	r := router.New()

	api := r.Group("/api")

	ws := store.NewWikiStore("Joinc")
	h := handler.NewHandler(ws)
	h.Register(api)
	fmt.Println(api)
	r.Logger.Fatal(r.Start(":8888"))
}
