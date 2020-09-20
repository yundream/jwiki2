package main

import (
	"fmt"
	"os"

	"joinc.co.kr/jwiki/db"
	"joinc.co.kr/jwiki/handler"
	"joinc.co.kr/jwiki/middleware"
	"joinc.co.kr/jwiki/router"
	"joinc.co.kr/jwiki/store"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := router.New()
	d, err := db.New()
	if err != nil {
		fmt.Println("DB Error", err.Error())
		os.Exit(1)
	}
	db.AutoMigration(d)

	api := r.Group("/api", middleware.ServerMiddleware)

	ws := store.NewWikiStore(d)
	h := handler.NewHandler(ws)
	h.Register(api)
	fmt.Println(api)
	r.Logger.Fatal(r.Start(":8888"))
}
