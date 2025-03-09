package helpers

import (
	"database/sql"
	"log"
	"time"
)

func ServerRoutine() {
	go func() {
		for {
			time.Sleep(time.Minute)

			err := delExpiredSessions(DataBase)
			if err != nil {
				log.Printf("Error cleaning up expired tokens: %v", err)
			} else {
				log.Println("Expired tokens cleaned up successfully.")
			}
		}
	}()
}

func delExpiredSessions(db *sql.DB) error {
	query := `DELETE FROM sessions WHERE expires_at < CURRENT_TIMESTAMP`
	_, err := db.Exec(query)
	return err
}
