package lab

import (
	"database/sql"
	"log"
	"time"
)

func ServerRoutine(db *sql.DB) {
	go func() {
		for {
			time.Sleep(time.Minute)

			// Run cleanup function
			err := cleanupExpiredsessions(db)
			if err != nil {
				log.Printf("Error cleaning up expired tokens: %v", err)
			} else {
				log.Println("Expired tokens cleaned up successfully.")
			}
		}
	}()
}

func cleanupExpiredsessions(db *sql.DB) error {
	query := `DELETE FROM sessions WHERE expires_at < CURRENT_TIMESTAMP`
	_, err := db.Exec(query)
	return err
}
