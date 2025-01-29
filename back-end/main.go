package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var HtmlTemplates *template.Template

func init() {
	var err error
	HtmlTemplates, err = template.ParseGlob("./front-end/templates/*.html")
	if err != nil {
		fmt.Println("Error parsing templates: ", err.Error())
		//! send internal server error here instead of quitting
		os.Exit(1)
	}
}

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
