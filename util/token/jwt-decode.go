package token

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
