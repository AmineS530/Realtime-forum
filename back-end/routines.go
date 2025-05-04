package helpers

import (
	"database/sql"
	"time"
)

func ServerRoutine() {
	go func() {
		for {
			time.Sleep(time.Minute)

			// Run cleanup function
			delExpiredSessions(DataBase)
		}
	}()
}

func delExpiredSessions(db *sql.DB) error {
	query := `DELETE FROM sessions WHERE expires_at < CURRENT_TIMESTAMP`
	_, err := db.Exec(query)
	if err != nil {
		ErrorLog.Printf("Error cleaning up expired tokens: %v", err)
	} else {
		InfoLog.Println("Expired tokens cleaned up successfully.")
	}
	return err
}
