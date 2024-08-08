package main

import (
	"fmt"
	"net/http"
	handler "proxyvercel/api"
)

func main() {
	http.HandleFunc("/webhook", handler.HandleWebhook)

	port := ":8443"
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
