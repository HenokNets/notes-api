package main

import (
	"fmt"
	"log"
	"net/http"
	"notes-api/internal/server"
)

func main() {
	handler := server.NewServer()
	port := ":8080"
	fmt.Println("Starting the server")
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
