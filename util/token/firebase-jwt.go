package token

import "github.com/k-washi/jwt-decode/jwtdecode"

func DecodeToken(jwt string) (*jwtdecode.FireBaseCustomToken, error) {
	hCS, err := jwtdecode.JwtDecode.DecomposeFB(jwt)
	if err != nil {
		return nil, err
	}
	payload, err := jwtdecode.JwtDecode.DecodeClaimFB(hCS[1])
	if err != nil {
		return nil, err
	}
	return payload, nil
}
