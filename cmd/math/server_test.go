package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestMathServer(t *testing.T) {
	t.Run("add numbers in url parameters", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/add?num=4&num=5&num=32", nil)
		request.Header.Add("Authorization", "Bearer SUPER_SECRET_API_KEY")

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
		request.Header.Add("Authorization", "Bearer SUPER_SECRET_API_KEY")

		MathServer(response, request)

		got := response.Body.String()
		want := "41"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("add numbers in json", func(t *testing.T) {
		data := []byte(`{
    "nums": [4, 5, 32]
}`)

		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(data))
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Authorization", "Bearer SUPER_SECRET_API_KEY")

		MathServer(response, request)

		got := response.Body.String()
		want := "41"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should return a 401 response code if no auth token is provided", func(t *testing.T) {
		data := []byte(`{
    "nums": [4, 5, 32]
}`)

		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(data))
		request.Header.Add("Content-Type", "application/json")

		MathServer(response, request)

		got := response.Code
		want := http.StatusUnauthorized

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})


}
