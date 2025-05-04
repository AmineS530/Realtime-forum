package database

import (
	"database/sql"
	"log"
	"os"

	helpers "RTF/back-end"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// Open log files
	logFile, err := os.OpenFile("./ServerLogs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o664)
	if err != nil {
		log.Fatalf("Failed to open app log file: %v", err)
	}

	errorFile, err := os.OpenFile("./ServerLogs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o664)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}

	// Setup loggers
	helpers.InfoLog = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	helpers.ErrorLog = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Set default log output for Fatal logs
	log.SetOutput(errorFile)
}

func SetTables() *sql.DB {
	db, err := sql.Open("sqlite3", "./DataBase/RTF.db")
	if err != nil {
		helpers.ErrorLog.Fatalf("Error opening database: %v", err)
	}

	sqlContent, err := os.ReadFile("./DataBase/schema.sql")
	if err != nil {
		helpers.ErrorLog.Fatalf("Error reading SQL file: %v", err)
	}

	_, err = db.Exec(string(sqlContent))
	if err != nil {
		helpers.ErrorLog.Fatalf("Error executing SQL: %v", err)
	}
	helpers.InfoLog.Println("Database successfully created!")
	return db
}
