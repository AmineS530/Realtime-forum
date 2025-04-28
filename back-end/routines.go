package helpers

import (
	"database/sql"
	"time"

	"RTF/global"
)

func ServerRoutine() {
	go func() {
		for {
			time.Sleep(time.Minute)

			// Run cleanup function
			delExpiredSessions(global.DataBase)
		}
	}()
}

func delExpiredSessions(db *sql.DB) error {
	query := `DELETE FROM sessions WHERE expires_at < CURRENT_TIMESTAMP`
	_, err := db.Exec(query)
	if err != nil {
		global.ErrorLog.Printf("Error cleaning up expired tokens: %v", err)
	} else {
		global.InfoLog.Println("Expired tokens cleaned up successfully.")
	}
	return err
}
