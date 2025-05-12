package lab

import (
	"encoding/json"
	"log"
)

func GenerateNotification(message, notificationType string) string {
	// Create notification struct
	notification := Notification{
		Message: message,
		Type:    notificationType,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(notification)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return `{"message": "Error generating notification", "type": "error"}`
	}

	return string(jsonData)
}

type Notification struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
