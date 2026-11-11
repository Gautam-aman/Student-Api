package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gautam-aman/student-api/cmd/internal/config"
)

func main() {

	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Println("Server started")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Failed to start")
	}

}
