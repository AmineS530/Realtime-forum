package main

import (
	"fmt"
	"net/http"
	"os"

	database "RTF/DataBase"
	helpers "RTF/back-end"
	handlers "RTF/back-end/goFiles"
)

func init() {
	helpers.DataBase = database.SetTables()
}

func main() {
	defer helpers.DataBase.Close()
	httpPort := os.Getenv("PORT")

	if httpPort == "" {
		httpPort = "9090"
	}

	srv := http.Server{
		Addr:    ":" + httpPort,
		Handler: handlers.Routes(),
	}

	fmt.Printf("Starting server on: \033[1;32mhttp://localhost%s\033[0m\n", srv.Addr)
	// Start the server
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Unable to start server. Reason: %v\n", err)
	}
}
