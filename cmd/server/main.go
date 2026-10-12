package main

import (
	"log"
	"net/http"
	"os"

	"mathmatic-test-server/internal/di"
)

// @title       mathmatic-test-server API
// @version     1.0
// @description 수학 문제 출제 및 채점 서버 API
// @BasePath    /api
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
