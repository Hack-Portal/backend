package jwt

type jwtDecodeInterface interface {
	DecomposeFB(string) ([]string, error)
	DecodeClaimFB(string) (*FireBaseCustomToken, error)
}

var (
	//JwtDecode jwt decode and parse user info in firebase
	JwtDecode jwtDecodeInterface
)

type jwtDecode struct{}

func init() {
	JwtDecode = &jwtDecode{}
}

func ValidJWTtoken(token string) (*FireBaseCustomToken, error) {
	hCS, err := JwtDecode.DecomposeFB(token)
	if err != nil {
		return nil, err
	}

	payload, err := JwtDecode.DecodeClaimFB(hCS[1])
	if err != nil {
		return nil, err
	}

	// if payload.Expires-time.Now().Unix() < 0 {
	// 	return nil, errors.New("トークンの期限きれてんで？")
	// }

	return payload, nil
}
