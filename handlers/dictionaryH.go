package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	dictionary "github.com/timoteoBone/dictionaryAPI/data"
)

type Translation struct {
	l *log.Logger
}

func NewTranslation(l *log.Logger) *Translation {
	return &Translation{l}
}

func (t *Translation) GetTranslation(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	t.l.Println(dictionary.Translation(vars["id"]))

}
