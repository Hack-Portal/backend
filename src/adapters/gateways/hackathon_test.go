package gateways

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lib/pq"
)

func TestHacahaton(t *testing.T) {
	t.Run("Create", testCreateHackathon)
}

func testCreateHackathon(t *testing.T) {
	testCases := []struct {
		name    string
		arg     *models.Hackathon
		wantErr error
		errCode string
	}{
		{
			name: "success case",
			arg: &models.Hackathon{
				HackathonID: "hackathon_id",
				Name:        "Hackathon",
				Icon:        "https://icon.com",
				Link:        "https://link.com",
				Expired:     time.Now().AddDate(0, 0, 10),
				StartDate:   time.Now(),
				Term:        1,

				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
		},
		{
			name: "failed case",
			arg: &models.Hackathon{
				HackathonID: "hackathon_id",
				Name:        "Hackathon",
				Icon:        "https://icon.com",
				Link:        "https://link.com",
				Expired:     time.Now().AddDate(0, 0, 10),
				StartDate:   time.Now(),
				Term:        1,

				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			// 重複エラー
			wantErr: &pq.Error{Code: pgerrcode.DuplicateColumn},
			errCode: pgerrcode.DuplicateColumn,
		},
	}
	hg := NewHackathonGateway(dbconn, nil)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := hg.Create(context.Background(), tc.arg, []int64{1, 2, 3})
			if err != nil {
				if reflect.TypeOf(err) != reflect.TypeOf(tc.wantErr) {
					t.Errorf("got %v, want %s", reflect.TypeOf(err), reflect.TypeOf(tc.wantErr))
				} else {
					var pgErr *pgconn.PgError
					if errors.As(err, &pgErr) {
						if pgErr.Code != tc.errCode {
							t.Errorf("got %v, want %v", pgErr.Code, tc.errCode)
						}
					}
				}
			}
		})
	}
}
