package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	n := Controller{}
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	n.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatalf("wrong code returned: %d", w.Code)
	}

	body := w.Body.String()
	if body != fmt.Sprintf("Hi there, I love !") {
		t.Fatalf("wrong body returned: %s", body)
	}
}
