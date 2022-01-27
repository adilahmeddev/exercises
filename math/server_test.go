package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestMathServer(t *testing.T) {
	t.Run("add numbers in url parameters", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/add?num=4&num=5&num=32", nil)
		response := httptest.NewRecorder()

		MathServer(response, request)

		got := response.Body.String()
		want := "41"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("add numbers in form encoded body", func(t *testing.T) {
		data := url.Values{}
		data.Add("num", "4")
		data.Add("num", "5")
		data.Add("num", "32")

		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/add", strings.NewReader(data.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		MathServer(response, request)

		got := response.Body.String()
		want := "41"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

