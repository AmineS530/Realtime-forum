package helpers

import (
	"database/sql"
	"log"
)

var (
	DataBase *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

type Placeholder struct {
	Online bool
}
