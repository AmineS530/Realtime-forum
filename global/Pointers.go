package global

import (
	"database/sql"
	"html/template"
	"log"
)

var (
	DataBase *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger

	HtmlTemplates *template.Template
)
