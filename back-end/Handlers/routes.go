package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	helpers "RTF/back-end"
)

var (
	HtmlTemplates *template.Template
	Err           error
)

type ErrorPage struct {
	Num int
	Msg string
}

func init() {
	HtmlTemplates, Err = template.ParseGlob("./front-end/templates/*.html")
	if Err != nil {
		fmt.Println("Error parsing templates: ", Err.Error())
		os.Exit(1)
	}
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	for _, middleware := range Middleware {
		middleware(IndexHandler)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if Err != nil {
			fmt.Println("Error parsing templates: ", Err.Error())
			ErrorPagehandler(w, http.StatusInternalServerError)
			return
		}
		IndexHandler(w, r)
	})
	mux.Handle("/front-end/styles/", http.StripPrefix("/front-end/styles/", http.FileServer(http.Dir("./front-end/styles"))))
	mux.Handle("/front-end/scripts/", http.StripPrefix("/front-end/scripts/", http.FileServer(http.Dir("./front-end/scripts"))))
	mux.HandleFunc("/api/v1/get/{type}", dumbjson)
	mux.HandleFunc("/api/v1/ws", handleConnections)
	mux.HandleFunc("/api/login", LoginHandler)
	mux.HandleFunc("/api/register", RegisterHandler)
	return mux
}

// TODO sMArT
func dumbjson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	x := r.PathValue("type")

	// ErrorPagehandler(w, http.StatusInternalServerError, "azerqsdfwxcv")
	if x == "comments" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(commentjson))
	} else if x == "posts" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(postjson))
	} else if x == "history" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(historyjson))
	} else {
		ErrorPagehandler(w, http.StatusNotFound)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if Err != nil {
		ErrorPagehandler(w, http.StatusInternalServerError)
		return
	}
	if r.URL.Path == "/" {
		if err := HtmlTemplates.ExecuteTemplate(w, "index.html", helpers.Placeholder{
			Online: false,
		}); err != nil {
			fmt.Println("Error executing template: ", err.Error())
			ErrorPagehandler(w, http.StatusInternalServerError)
			return
		}
	} else {
		ErrorPagehandler(w, http.StatusNotFound)
		return
	}
}

func ErrorPagehandler(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	errorData := ErrorPage{
		Num: statusCode,
		Msg: http.StatusText(statusCode),
	}

	if err := HtmlTemplates.ExecuteTemplate(w, "error_page.html", errorData); err != nil {
		//	fmt.Println("Error executing template: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}