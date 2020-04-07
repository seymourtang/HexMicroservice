package api

import (
	"HexMicroservice/shortener"
	"net/http"
)

type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}
type handler struct {
	redirectService shortener.RedirectService
}
