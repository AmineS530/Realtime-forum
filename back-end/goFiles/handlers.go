package handlers

import (
	"fmt"
	"net/http"
	"strings"

	helpers "RTF/back-end"
	"RTF/back-end/goFiles/auth"
	"RTF/global"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if helpers.Err != nil {
		auth.JsRespond(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		auth.JsRespond(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/api/") {
		auth.JsRespond(w, "API endpoint not found", http.StatusNotFound)
		return
	}
	if r.URL.Path == "/" || r.URL.Path == "/profile" {
		if err := global.HtmlTemplates.ExecuteTemplate(w, "index.html", nil); err != nil {
			fmt.Println("Error executing template: ", err.Error())
			auth.JsRespond(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	} else {
		// helpers.ErrorPagehandler(w, http.StatusNotFound)
		auth.JsRespond(w, "Page not found", http.StatusNotFound)
		return
	}
}
