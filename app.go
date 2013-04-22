package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "User-Agent: %s\n", r.Header.Get("User-Agent"))
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Current time: %s\n", time.Now().String())
}

func main() {
	var addr = flag.String("http", ":8080", "HTTP service address")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("Starting HTTP server on", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
