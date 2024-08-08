package main

import (
	"fmt"
	"net/http"

	handler "github.com/huof6829/proxyvercel/api"
)

func main() {
	http.HandleFunc("/api/webhook", handler.HandleWebhook)

	port := ":8441"
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
