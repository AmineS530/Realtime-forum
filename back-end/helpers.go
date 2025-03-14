package helpers

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	DataBase *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

type ErrorPage struct {
	Num int
	Msg string
}

type Placeholder struct {
	Online bool
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
