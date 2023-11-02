package gateways

import (
	"testing"
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/utils"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
)

func TestCreate(t *testing.T) {
	hg := NewHackathonGateway(db)

	testCases := []struct {
		name string
		arg  params.HackathonCreate
		err  error
	}{
		{
			name: "正常系",
			arg: params.HackathonCreate{
				Hackathon: entities.Hackathon{
					HackathonID: utils.NewUUID(),
					Name:        utils.RandomString(10),
					Icon:        utils.RandomString(10),
					Link:        utils.RandomString(10),
					StartDate:   time.Now(),
					Term:        1,
					Expired:     time.Now().AddDate(0, 0, 1),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					IsDelete:    false,
				},
				Statuses: []int32{1, 2},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := hg.Create(tc.arg)
			if err != tc.err {
				t.Errorf("got: %v, want: %v", err, tc.err)
			}
		})
	}
}

func addTestData(hg dai.HackathonRepository, t *testing.T) params.HackathonCreate {
	arg := params.HackathonCreate{
		Hackathon: entities.Hackathon{
			HackathonID: utils.NewUUID(),
			Name:        utils.RandomString(10),
			Icon:        utils.RandomString(10),
			Link:        utils.RandomString(10),
			StartDate:   time.Now(),
			Term:        1,
			Expired:     time.Now().Add(time.Duration(5 * time.Minute)),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			IsDelete:    false,
		},
		Statuses: utils.RandomIntArr(2, utils.RandomInt(2)),
	}
	if err := hg.Create(arg); err != nil {
		t.Fatalf("Cannnot Create Hackathon :%v", err)
	}
	return arg
}

func TestReadAll(t *testing.T) {
	hg := NewHackathonGateway(db)

	var hackathons []params.HackathonCreate

	for i := 0; i < 6; i++ {
		hackathons = append(hackathons, addTestData(hg, t))
	}

	testCases := []struct {
		name    string
		arg     params.HackathonReadAll
		err     error
		wantLen int
	}{
		{
			name: "正常系(タグでソートする場合)",
			arg: params.HackathonReadAll{
				Limit:   3,
				Offset:  0,
				SortTag: []int32{2},
			},
			err:     nil,
			wantLen: 3,
		},
		{
			name: "正常系(タグでソートしないパターン)",
			arg: params.HackathonReadAll{
				Limit:  3,
				Offset: 0,
			},
			err:     nil,
			wantLen: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hackatons, _, err := hg.ReadAll(tc.arg)
			if err != tc.err {
				t.Fatalf("got: %v, want: %v", err, tc.err)
			}
			
			if len(hackatons) != tc.wantLen {
				t.Fatalf("got: %v, want: %v", len(hackatons), tc.wantLen)
			}
		})
	}
}
