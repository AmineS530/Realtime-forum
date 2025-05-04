package jwt

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strings"
	"time"

	helpers "RTF/back-end"
)

var Time_to_Expire = time.Hour * 24

type JwtPayload struct {
	Sub      int    `json:"sub,string"`
	Username string `json:"username"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}

// LoadSecret manually reads the .env file and retrieves SECRET_KEY
func LoadSecret() string {
	file, err := os.Open(".env")
	if err != nil {
		helpers.ErrorLog.Fatal("Error loading .env file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "SECRET_KEY=") {
			return strings.TrimPrefix(line, "SECRET_KEY=")
		}
	}

	helpers.ErrorLog.Fatal("SECRET_KEY not found in .env file")
	return ""
}

// base64Encode encodes data to a URL-safe base64 string
func base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// base64Decode decodes a URL-safe base64 string
func base64Decode(encoded string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(encoded)
}

func signMessage(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	signature := h.Sum(nil)
	return base64Encode(signature)
}
