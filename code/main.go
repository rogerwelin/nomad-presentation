package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Hello struct {
	Msg string `json:"msg"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := Hello{
		Msg: "hello meetup",
	}
	json, _ := json.Marshal(resp)
	fmt.Fprintln(w, string(json))
}

func main() {
	addr := os.Getenv("NOMAD_ADDR_http")
	if addr == "" {
		addr = "0.0.0.0:5000"
	}

	router := http.NewServeMux()
	router.HandleFunc("/hello", helloHandler)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Handler:      router,
	}
	log.Fatal(srv.ListenAndServe())
}
