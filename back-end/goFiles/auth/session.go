package auth

import (
	"time"

	helpers "RTF/back-end"
	jwt "RTF/back-end/goFiles/JWT"

	"github.com/gofrs/uuid"
)

// Todo: send errors in js notifications
// todo: check for session id

func CheckSession(userID int, username string) (string, string, error) {
	activeSessionID, _ := CheckActiveSession(userID)
	// sessions := getElemVal("session_id", "sessions", `id = "`+strconv.Itoa(userID)+`"`)
	// sessionList, ok := sessions.([]interface{})
	// if ok && len(sessionList) > 0 {
	// 	// Invalidate all previous sessions
	// 	for _, s := range sessionList {
	// 		if _, valid := s.(string); valid {
	// 			InvalidateSession(userID)
	// 		}
	// 	}
	// }
	if activeSessionID != "" {
		InvalidateSession(userID)
	}

	// Create a new session
	sessionID, err := createSession(userID)
	if err != nil {
		return "", "", err
	}

	payload := jwt.CreateJwtPayload(userID, username)
	jwtToken := jwt.Generate(payload)

	return jwtToken, sessionID, nil
}

func createSession(userID int) (string, error) {
	// Generate unique session ID (UUID or random string)
	sessionID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	// Set session expiration time from int64
	expiresAt := time.Now().Add(time.Duration(jwt.Time_to_Expire))

	// Insert the session into the database
	_, err = helpers.DataBase.Exec(`
        INSERT INTO sessions (user_id, session_id , expires_at) 
        VALUES (?, ?, ?)
    `, userID, sessionID.String(), expiresAt)
	if err != nil {
		return "", err
	}

	return sessionID.String(), nil
}

func CheckActiveSession(userID int) (string, error) {
	var sessionID string
	err := helpers.DataBase.QueryRow(`
        SELECT session_id 
        FROM sessions 
        WHERE user_id = ? AND expires_at > CURRENT_TIMESTAMP
    `, userID).Scan(&sessionID)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

// logout
// todo merge with other logout
func InvalidateSession(userID int) error {
	_, err := helpers.DataBase.Exec(`
        DELETE FROM sessions 
        WHERE user_id = ?
    `, userID)
	return err
}
