package gateways

import (
	"log"
	"testing"
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/utils"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
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

func TestReadAll(t *testing.T) {
	hg := NewHackathonGateway(db)

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
			hackatons, err := hg.ReadAll(tc.arg)
			if err != tc.err {
				t.Fatalf("got: %v, want: %v", err, tc.err)
			}
			log.Println(hackatons)
			if len(hackatons) != tc.wantLen {
				t.Fatalf("got: %v, want: %v", len(hackatons), tc.wantLen)
			}
		})
	}
}
