package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	helpers "RTF/back-end"
	"RTF/back-end/goFiles/auth"
	"RTF/back-end/goFiles/ws"
)

var (
	HtmlTemplates *template.Template
	Err           error
)

func init() {
	HtmlTemplates, Err = template.ParseGlob("./front-end/templates/*.html")
	if Err != nil {
		fmt.Println("Error parsing templates: ", Err.Error())
		os.Exit(1)
	}
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	protectedRoutes := []string{"/api/v1/get/{type}"} // Add more as needed

	for _, route := range protectedRoutes {
		handler := dumbjson
		for _, middleware := range auth.Middleware {
			handler = middleware(handler)
		}
		mux.HandleFunc(route, handler)
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
	mux.Handle("/front-end/images/", http.StripPrefix("/front-end/images/", http.FileServer(http.Dir("./front-end/images"))))

	mux.HandleFunc("/api/ws", ws.HandleWebSocket)

	mux.HandleFunc("/api/check-auth", auth.CheckAuthHandler)
	mux.HandleFunc("/api/login", auth.LoginHandler)
	mux.HandleFunc("/api/register", auth.RegisterHandler)
	mux.HandleFunc("/api/logout", auth.Logout)
	helpers.ServerRoutine()

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
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path == "/" {
		if err := HtmlTemplates.ExecuteTemplate(w, "index.html", nil); err != nil {
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
	errorData := helpers.ErrorPage{
		Num: statusCode,
		Msg: http.StatusText(statusCode),
	}

	if err := HtmlTemplates.ExecuteTemplate(w, "error_page.html", errorData); err != nil {
		//	fmt.Println("Error executing template: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
