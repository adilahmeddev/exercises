package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main(){
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--web-server" {
		r := mux.NewRouter()
		r.HandleFunc("/add", MathServer)
		r.HandleFunc("/fibonacci/{num}", FibonacciServer)

		server := &http.Server{
			Handler: r,
			Addr: "127.0.0.1:5000",
		}
		log.Fatal(server.ListenAndServe())
	}
}

