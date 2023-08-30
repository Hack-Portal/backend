package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"firebase.google.com/go/auth"
)

//FireBaseCustomToken auth.tokenをemailで拡張
/*
   type Token struct{
       Issuer string "json:\"iss\"";
       Audience string "json:\"aud\"";
       Expires int64 "json:\"exp\"";
       IssuedAt int64 "json:\"iat\"";
       Subject string "json:\"sub,omitempty\"";
       UID string "json:\"uid,omitempty\"";
       Claims map[string]interface{}
       "json:\"-\""
   }
*/

type FireBaseCustomToken struct {
	auth.Token
	Email string `json:"email"`
}

// DecomposeFB  JWTをHeader, claims, 署名に分解
func (s *jwtDecode) DecomposeFB(jwt string) ([]string, error) {
	hCS := strings.Split(jwt, ".")
	if len(hCS) == 3 {
		return hCS, nil
	}
	return nil, errors.New("Error jwt str decompose: inccorrect number of segments")

}

// DecodeClaimFB JWTのclaims部分をFireBaseCustomTokenの構造体にデコード
func (s *jwtDecode) DecodeClaimFB(payload string) (*FireBaseCustomToken, error) {
	payloadByte, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil, errors.New("Error jwt token decode: " + err.Error())
	}

	var tokenJSON FireBaseCustomToken
	err = json.Unmarshal(payloadByte, &tokenJSON)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, errors.New("Error jwt token unmarshal: " + err.Error())
	}

	return &tokenJSON, nil
}
