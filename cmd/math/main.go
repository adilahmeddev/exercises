package main

import (
	"log"
	"net/http"
	"os"
)

func main(){
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--web-server" {
		handler := http.HandlerFunc(MathServer)
		log.Fatal(http.ListenAndServe(":5000", handler))
	}
}