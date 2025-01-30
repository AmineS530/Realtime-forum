package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var (
	HtmlTemplates *template.Template
	Err           error
)

type ErrorPage struct {
	Num string
	Msg string
}

func init() {
	var err error
	HtmlTemplates, Err = template.ParseGlob("./front-end/templates/*.html")
	if err != nil {
		fmt.Println("Error parsing templates: ", err.Error())
		//! send internal server error here instead of quitting
		os.Exit(1)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if Err != nil {
		ErrorPagehandler(w, http.StatusInternalServerError, Err.Error())
		return
	}
	if r.URL.Path == "/" {
		if err := HtmlTemplates.ExecuteTemplate(w, "first.html", nil); err != nil {
			ErrorPagehandler(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		ErrorPagehandler(w, http.StatusNotFound, "Where do you think you're going ehehehe?")
		return
	}
}

func ErrorPagehandler(w http.ResponseWriter, statusCode int, errMsg string) {
	w.WriteHeader(statusCode)
	errorData := ErrorPage{
		Num: http.StatusText(statusCode),
		Msg: errMsg,
	}

	if err := HtmlTemplates.ExecuteTemplate(w, "Error.html", errorData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
