package main

import (
	"github.com/FernandoCagale/c4-order/api/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	app, e := SetupApplication()

	if e != nil {
		panic("Erro to start application")
	}

	app.MakeEvents()

	router := app.MakeHandlers()

	router.Use(middleware.Header)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
