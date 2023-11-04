package usecases

import (
	"testing"
	"time"

	"github.com/hackhack-Geek-vol6/backend/src/adapters/gateways/mock"
	"github.com/hackhack-Geek-vol6/backend/src/adapters/presenters"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/ports"
)

func TestCreate(t *testing.T) {

	hu := NewHackathonInteractor(
		presenters.NewHackathonOutputBoundary(),
		mock.NewMockHackathonRepository(),
		mock.NewMockFirebaseRepository(),
	)

	testCases := []struct {
		name       string
		arg        input.HackathonCreate
		image      []byte
		wantStatus int
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
			image:      []byte("test"),
			wantStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, _ := hu.Create(tc.arg, tc.image)
			if status != tc.wantStatus {
				t.Errorf("want: %v, got: %v", tc.wantStatus, status)
			}
		})
	}
}

func addTestData(t *testing.T, hu ports.HackathonInputBoundary) {
	arg := input.HackathonCreate{
		Name:       "test",
		Link:       "test",
		Expired:    time.Now().AddDate(0, 0, 1),
		StartDate:  time.Now(),
		Term:       1,
		StatusTags: "1,2",
	}
	image := []byte("test")

	hu.Create(arg, image)
}

func TestReadAll(t *testing.T) {
	hu := NewHackathonInteractor(
		presenters.NewHackathonOutputBoundary(),
		mock.NewMockHackathonRepository(),
		mock.NewMockFirebaseRepository(),
	)

	// var args params.HackathonCreate
	for i := 0; i < 10; i++ {
		addTestData(t, hu)
	}

	testCases := []struct {
		name       string
		arg        input.HackathonReadAll
		wantLen    int
		wantStatus int
	}{
		{
			name: "success",
			arg: input.HackathonReadAll{
				PageSize: 3,
				PageID:   1,
				SortTag:  []int32{1},
			},
			wantLen:    3,
			wantStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, hackathons := hu.ReadAll(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("want: %v, got: %v", tc.wantStatus, status)
			}

			if len(hackathons) != tc.wantLen {
				t.Errorf("want: %v, got: %v", tc.wantLen, len(hackathons))
			}

			for _, hackathon := range hackathons {
				if containsArr(hackathon.StatusTags, tc.arg.SortTag) {
					t.Errorf("want: %v, got: %v", tc.arg.SortTag, hackathon.StatusTags)
				}
			}
		})
	}
}

func containsArr(value []entities.StatusTag, target []int32) bool {
	for _, t := range target {
		for _, v := range value {
			if t == v.StatusID {
				return true
			}
		}
	}
	return false
}
