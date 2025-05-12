package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	DataBase      *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Err           error
	HtmlTemplates *template.Template
)

func init() {
	HtmlTemplates, Err = template.ParseGlob("./front-end/templates/*.html")
	if Err != nil {
		fmt.Println("Error parsing templates: ", Err.Error())
		os.Exit(1)
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
type ErrorPage struct {
	Num int
	Msg string
}

func EntryExists(elem, value, from string, checkLower bool) (int, bool) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", from, elem)
	if checkLower {
		query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE LOWER(%s) = LOWER(?)", from, elem)
	}

	err := DataBase.QueryRow(query, value).Scan(&count)
	if err != nil {
		ErrorLog.Fatalln("Database error:", err)
		return -1, false
	}

	return count, count > 0
}

func JsRespond(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
		Code:  code,
	})
}
