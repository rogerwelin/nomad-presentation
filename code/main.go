package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home\n")
}

func main() {
	addr := os.Getenv("NOMAD_ADDR_http")
	if addr == "" {
		addr = "0.0.0.0:5000"
	}

	router := http.NewServeMux()
	router.HandleFunc("/home", rootHandler)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Handler:      router,
	}
	log.Fatal(srv.ListenAndServe())
}
