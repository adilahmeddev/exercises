package main

import (
	"bytes"
	"github.com/matryer/is"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("three number arguments", func(t *testing.T) {
		is := is.New(t)
		output := bytes.Buffer{}
		args := []string{"10", "5", "7"}

		run(args, &output)

		want:= "22"
		got := output.String()

		is.Equal(want, got)
	})

	t.Run("invalid arguments", func(t *testing.T) {
		is := is.New(t)
		output := bytes.Buffer{}
		args := []string{"adil", "5", "7"}

		run(args, &output)

		want:= "invalid input"
		got := output.String()

		is.Equal(want, got)
	})
}
