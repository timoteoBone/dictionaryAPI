package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/timoteoBone/dictionaryAPI/handlers"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	dh := handlers.NewTranslation(l)

	rt := mux.NewRouter()

	getRouter := rt.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/translate/{id}", dh.GetTranslation)

	putRouter := rt.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(("/add/{id}/{translation}"), dh.AddTranslation)

	s := &http.Server{
		Handler:        rt,
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
