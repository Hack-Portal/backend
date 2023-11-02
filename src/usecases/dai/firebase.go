package dai

type FirebaseRepository interface {
	UploadFile(string, []byte) (string, error)
	DeleteFile(string) error
}
