package mock

import "github.com/hackhack-Geek-vol6/backend/src/usecases/dai"

type MockFirebaseRepository struct {
	image map[string][]byte
}

func NewMockFirebaseRepository() dai.FirebaseRepository {
	return &MockFirebaseRepository{}
}

func (m *MockFirebaseRepository) UploadFile(hackathonID string, image []byte) (string, error) {
	i := map[string][]byte{
		hackathonID: image,
	}
	m.image = i
	return hackathonID + ".jpg", nil
}

func (m *MockFirebaseRepository) DeleteFile(string) error {
	return nil
}
