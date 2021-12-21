package handlers

import (
	"log"
	"net/http"
)

type Translation struct {
	l *log.Logger
}

func NewTranslation(l *log.Logger) *Translation {
	return &Translation{l}
}

func (t *Translation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
