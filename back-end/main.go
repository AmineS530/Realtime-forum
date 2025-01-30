package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	httpPort := os.Getenv("PORT")

	if httpPort == "" {
		httpPort = "9090"
	}

	srv := http.Server{
		Addr:    ":" + httpPort,
		Handler: http.DefaultServeMux,
	}

	fmt.Printf("Starting server on: \033[1;32mhttp://localhost%s\033[0m\n", srv.Addr)
	// Start the server
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Unable to start server. Reason: %v\n", err)
	}
}
