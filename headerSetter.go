package main

import (
	"net/http"
)

// HeaderSetter struct
type HeaderSetter struct {
	key, val string
	handler  http.Handler
}

func NewHeaderSetter(key, val string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return HeaderSetter{key, val, h}
	}
}

func (hs HeaderSetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(hs.key, hs.val)
	hs.handler.ServeHTTP(w, r)
}
