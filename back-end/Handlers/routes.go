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
	HtmlTemplates, Err = template.ParseGlob("./front-end/templates/*.html")
	if Err != nil {
		fmt.Println("Error parsing templates: ", Err.Error())
		//! send internal server error here instead of quitting
		os.Exit(1)
	}
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		IndexHandler(w, r)
	})
	mux.HandleFunc("/front-end/", func(w http.ResponseWriter, r *http.Request) {
		FileHandler(w, r)
	})
	return mux
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if Err != nil {
		ErrorPagehandler(w, http.StatusInternalServerError, Err.Error())
		return
	}
	if r.URL.Path == "/" {
		if err := HtmlTemplates.ExecuteTemplate(w, "index.html", nil); err != nil {
			ErrorPagehandler(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		ErrorPagehandler(w, http.StatusNotFound, "Where do you think you're going ehehehe?")
		return
	}
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if Err != nil {
		ErrorPagehandler(w, http.StatusInternalServerError, Err.Error())
		return
	}
	fmt.Println(r.URL.Path)
	http.ServeFile(w, r, r.URL.Path[1:])
}

func ErrorPagehandler(w http.ResponseWriter, statusCode int, errMsg string) {
	// w.WriteHeader(statusCode)
	http.Error(w, errMsg, statusCode)
}
