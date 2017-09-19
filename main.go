package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {

	h := http.NewServeMux()

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	c := &Controller{}
	h.Handle("/bar", c)

	chain := alice.New(
		NewHeaderSetter("X-FOO", "BAR"),
		NewHeaderSetter("X-BAZ", "BUZ"),
		logger,
	).Then(h)

	// hl := logger(h)
	// hhs := NewHeaderSetter("X-FOO", "BAR")(hl)

	err := http.ListenAndServe(":8080", chain)
	log.Fatal(err)

}
