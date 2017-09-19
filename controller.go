package main

import (
	"fmt"
	"net/http"
)

type Controller struct {
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
