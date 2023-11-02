package dai

type FirebaseRepository interface {
	UploadFile([]byte) (string, error)
	DeleteFile() error
}
