package usecases

import (
	"testing"
	"time"

	"github.com/hackhack-Geek-vol6/backend/src/adapters/gateways/mock"
	"github.com/hackhack-Geek-vol6/backend/src/adapters/presenters"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
)

func TestCreate(t *testing.T) {

	hu := NewHackathonInteractor(
		presenters.NewHackathonOutputBoundary(),
		mock.NewMockHackathonRepository(),
		mock.NewMockFirebaseRepository(),
	)

	testCases := []struct {
		name  string
		arg   input.HackathonCreate
		image []byte
		want  int

		buildStub func(*HackathonInteractor)
	}{
		{
			name: "success",
			arg: input.HackathonCreate{
				Name:       "test",
				Link:       "test",
				Expired:    time.Now().AddDate(0, 0, 1),
				StartDate:  time.Now(),
				Term:       1,
				StatusTags: "1,2",
			},
			image: []byte("test"),
			want:  200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, _ := hu.Create(tc.arg, tc.image)
			if status != tc.want {
				t.Errorf("want: %v, got: %v", tc.want, status)
			}
		})
	}
}
