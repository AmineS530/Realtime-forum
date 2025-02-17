package helpers

import "database/sql"

var DataBase *sql.DB

type Placeholder struct {
	Online bool
}
