package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Issued   string `json:"issued"`
	Audience string `json:"audience"`
	AuthTime int64  `json:"auth_time"`
	UserId   string `json:"user_id"`
	Subject  string `json:"subject"`
	IssuedAt int64  `json:"issued_at"`
	Expired  int64  `json:"expired"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// googleから公開鍵を取得する
func getGooglePublicKey() (map[string]interface{}, error) {
	response, err := http.Get("https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 公開鍵から解析する
func parsePublicKey(result, header map[string]interface{}) (*rsa.PublicKey, error) {
	kid := header["kid"].(string)
	certString := result[kid].(string)
	block, _ := pem.Decode([]byte(certString))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	return cert.PublicKey.(*rsa.PublicKey), nil
}

// jwtからヘッダを解析する
func parseHeader(tokenString string) (map[string]interface{}, error) {
	parts := strings.Split(tokenString, ".")
	headerJson, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}

	var header map[string]interface{}
	err = json.Unmarshal(headerJson, &header)
	if err != nil {
		return nil, err
	}
	return header, nil
}

// firebaseJWTがclaimを取得する
func CheckFirebaseJWT(tokenString string) (CustomClaims, error) {
	result, err := getGooglePublicKey()
	if err != nil {
		return CustomClaims{}, err
	}

	header, err := parseHeader(tokenString)
	if err != nil {
		return CustomClaims{}, err
	}

	rsaPublicKey, err := parsePublicKey(result, header)
	if err != nil {
		return CustomClaims{}, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return rsaPublicKey, nil
	})

	if err != nil {
		err = errors.New("failed to parse PEM block containing the public key")
		return CustomClaims{}, err
	}

	// Tokenを検証する
	claims, ok := token.Claims.(*CustomClaims)
	if !ok && token.Valid {
		err = errors.New("token is not valid")
		return CustomClaims{}, err
	}

	// 期限を検証する
	if time.Unix(claims.Expired, 0).Before(time.Now()) {
		err = errors.New("token is valid. but token is expired")
		return CustomClaims{}, err
	}

	return *claims, nil
}
