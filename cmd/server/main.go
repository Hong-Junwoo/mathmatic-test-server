package main

import (
	"log"
	"net/http"
	"os"

	"mathmatic-test-server/internal/di"
)

func main() {
	addr := ":8080"
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}

	log.Printf("서버 시작: %s", addr)
	if err := http.ListenAndServe(addr, di.NewHTTPHandler()); err != nil {
		log.Fatal(err)
	}
}
