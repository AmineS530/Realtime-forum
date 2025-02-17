package main

import (
	"fmt"
	"net/http"
	"os"

	handlers "RTF/back-end/Handlers"
)

func init() {
	// var err error
	// helpers.DataBase, err = sql.Open("sqlite3", "./forum.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// sqlFile := "./DataBase/schema.sql"
	// sqlContent, err := os.ReadFile(sqlFile)
	// if err != nil {
	// 	log.Fatal("Error at reading sql file!", err)
	// }

	// // Execute the SQL content to create tables.
	// _, err = helpers.DataBase.Exec(string(sqlContent))
	// if err != nil {
	// 	log.Fatal("Error at executing sql!", err)
	// }

	// fmt.Println("db successfully created!")
}

func main() {
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
