package main

import (
	"log"
	"net/http"
	"os"
)

func main(){
	args := os.Args[1:]
	if args[0] == "--web-server" {
		handler := http.HandlerFunc(MathServer)
		log.Fatal(http.ListenAndServe(":5000", handler))
	}
}