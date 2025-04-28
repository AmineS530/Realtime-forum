package jwt

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"RTF/structs"
)

var secretKey = LoadSecret()

func CreateJwtPayload(id int, username string) structs.JwtPayload {
	iat := time.Now().Unix()
	exp := iat + int64(Time_to_Expire.Seconds())

	return structs.JwtPayload{
		Sub:      id,
		Username: username,
		Iat:      iat,
		Exp:      exp,
	}
}

// Generate creates a JWT with a given payload
func Generate(payload structs.JwtPayload) string {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerJSON, _ := json.Marshal(header)
	payloadJSON, _ := json.Marshal(payload)

	encodedHeader := base64Encode(headerJSON)
	encodedPayload := base64Encode(payloadJSON)

	signature := signMessage(encodedHeader+"."+encodedPayload, secretKey)

	return encodedHeader + "." + encodedPayload + "." + signature
}

// Verify checks if the JWT signature is valid
func JWTVerify(token string) (*structs.JwtPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	decodedPayload, err := base64Decode(parts[1])
	if err != nil {
		return nil, errors.New("invalid payload encoding")
	}

	if signMessage(parts[0]+"."+parts[1], secretKey) != parts[2] {
		return nil, errors.New("invalid signature")
	}

	var payload structs.JwtPayload
	err = json.Unmarshal(decodedPayload, &payload)
	if err != nil {
		return nil, errors.New("invalid payload JSON")
	}

	return &payload, nil
}
