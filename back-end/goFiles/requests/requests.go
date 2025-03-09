package requests

import jwt "RTF/back-end/goFiles/JWT"

func HandleReq(jwt_token string) {
	payload, err := jwt.JWTVerify(jwt_token)
	if err != nil {
		// handle error
	}
	payload.Sub = 1
}
