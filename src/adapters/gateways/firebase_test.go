package gateways

import (
	"os"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/pkg/utils"
)

func loadTestImage() ([]byte, error) {
	return os.ReadFile("../../../color.png")
}

func TestFirebaseStorage(t *testing.T) {
	fb := NewFirebaseRepository(app)
	image, err := loadTestImage()
	if err != nil {
		t.Fatalf("failed to load image: %v", err)
	}

	testCases := []struct {
		name string
		id   string
		arg  []byte
		err  error
	}{
		{
			name: "正常系",
			id:   utils.NewUUID(),
			arg:  image,
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := fb.UploadFile(tc.id, tc.arg)
			if err != tc.err {
				t.Errorf("got: %v, want: %v", err, tc.err)
			}

			if err = fb.DeleteFile(tc.id); err != nil {
				t.Errorf("failed to delete file: %v", err)
			}
		})
	}
}
