package gateways

import (
	"context"
	"testing"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"gorm.io/gorm"
)

func TestStatusTag(t *testing.T) {
	t.Run("Create", testCreateStatusTag)
	t.Run("GetAll", TestGetAllStatusTag)
	t.Run("GetByID", TestGetStatusTagByID)
	t.Run("Update", TestUpdateStatusTag)
}

func testCreateStatusTag(t *testing.T) {
	testCases := []struct {
		name    string
		status  string
		wantErr error
	}{
		{
			name:    "正常系",
			status:  "test",
			wantErr: nil,
		},
	}

	sg := NewStatusTagGateway(dbconn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sg.Create(context.Background(), &models.StatusTag{
				Status: tc.status,
			})
			if err != tc.wantErr {
				t.Errorf("want: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}

func TestGetAllStatusTag(t *testing.T) {
	testCases := []struct {
		name           string
		statusesLength int
		wantErr        error
	}{
		{
			name:           "正常系",
			statusesLength: 1,
			wantErr:        nil,
		},
	}

	sg := NewStatusTagGateway(dbconn)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statuses, err := sg.FindAll(context.Background())
			if err != tc.wantErr {
				t.Errorf("want: %v, got: %v", tc.wantErr, err)
			}
			if len(statuses) != tc.statusesLength {
				t.Errorf("want: %v, got: %v", tc.statusesLength, len(statuses))
			}
		})
	}
}

func TestGetStatusTagByID(t *testing.T) {
	testCases := []struct {
		name        string
		id          int64
		wantErr     error
		wantErrCode int
	}{
		{
			name:    "正常系",
			id:      1,
			wantErr: nil,
		},
		{
			name:    "ID Not Found Error",
			id:      0,
			wantErr: gorm.ErrRecordNotFound,
		},
	}

	sg := NewStatusTagGateway(dbconn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sg.FindById(context.Background(), tc.id)
			if err != tc.wantErr {
				t.Errorf("want: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}

func TestUpdateStatusTag(t *testing.T) {
	// Updateのテストケースを並列に実行する
	testCases := []struct {
		name    string
		id      int64
		status  string
		wantErr error
	}{
		{
			name:    "正常系",
			id:      1,
			status:  "test1",
			wantErr: nil,
		},
		{
			name:    "ID Not Found Error",
			id:      0,
			status:  "test2",
			wantErr: gorm.ErrRecordNotFound,
		},
	}

	sg := NewStatusTagGateway(dbconn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sg.Update(context.Background(), &models.StatusTag{
				StatusID: tc.id,
				Status:   tc.status,
			})
			if err != tc.wantErr {
				t.Errorf("want: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}
