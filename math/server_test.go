package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMathServer(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/add?num=4&num=5&num=32", nil)
		response := httptest.NewRecorder()

		MathServer(response, request)

		got := response.Body.String()
		want := "41"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

