package interactors

import (
	"context"
	"testing"

	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/adapters/presenters"
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/stretchr/testify/require"
)

func TestStatusTag(t *testing.T) {
	t.Run("CreateStatusTag", testCreateStatusTag)
	t.Run("FindAllStatusTag", testFindAllStatusTag)
	t.Run("FindByIdStatusTag", testFindByIdStatusTag)
	t.Run("UpdateStatusTag", testUpdateStatusTag)
}

func testCreateStatusTag(t *testing.T) {
	testCases := []struct {
		name       string
		input      *request.CreateStatusTag
		wantStatus int
		want       *response.StatusTag
	}{
		{
			name: "success",
			input: &request.CreateStatusTag{
				Status: "test",
			},
			wantStatus: 201,
			want: &response.StatusTag{
				ID:     1,
				Status: "test",
			},
		},
		{
			name: "failed request field required",
			input: &request.CreateStatusTag{
				Status: "",
			},
			wantStatus: 400,
			want:       nil,
		},
	}
	si := NewStatusTagInteractor(gateways.NewStatusTagGateway(dbconn), presenters.NewStatusTagPresenter())
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, resp := si.CreateStatusTag(context.Background(), tc.input)
			if status != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, status)
			}
			require.Equal(t, resp, tc.want)
		})
	}
}

func testFindAllStatusTag(t *testing.T) {
	testCases := []struct {
		name       string
		wantStatus int
		want       []*response.StatusTag
	}{
		{
			name:       "success",
			wantStatus: 200,
			want: []*response.StatusTag{
				{
					ID:     1,
					Status: "test",
				},
			},
		},
	}

	si := NewStatusTagInteractor(gateways.NewStatusTagGateway(dbconn), presenters.NewStatusTagPresenter())
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, resp := si.FindAllStatusTag(context.Background())
			if status != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, status)
			}
			require.Equal(t, resp, tc.want)
		})
	}
}

func testFindByIdStatusTag(t *testing.T) {
	testCases := []struct {
		name       string
		input      *request.GetStatusTagByID
		wantStatus int
		want       *response.StatusTag
	}{
		{
			name: "success",
			input: &request.GetStatusTagByID{
				ID: 1,
			},
			wantStatus: 200,
			want: &response.StatusTag{
				ID:     1,
				Status: "test",
			},
		},
		{
			name: "failed request field required",
			input: &request.GetStatusTagByID{
				ID: 0,
			},
			wantStatus: 400,
			want:       nil,
		},
	}

	si := NewStatusTagInteractor(gateways.NewStatusTagGateway(dbconn), presenters.NewStatusTagPresenter())
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, resp := si.FindByIdStatusTag(context.Background(), tc.input)
			if status != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, status)
			}
			require.Equal(t, resp, tc.want)
		})
	}
}

func testUpdateStatusTag(t *testing.T) {
	testCases := []struct {
		name       string
		input      *request.UpdateStatusTag
		wantStatus int
		want       *response.StatusTag
	}{
		{
			name: "success",
			input: &request.UpdateStatusTag{
				ID:     1,
				Status: "test",
			},
			wantStatus: 200,
			want: &response.StatusTag{
				ID:     1,
				Status: "test",
			},
		},
		{
			name: "failed request field required",
			input: &request.UpdateStatusTag{
				ID:     1,
				Status: "",
			},
			wantStatus: 400,
			want:       nil,
		},
	}

	si := NewStatusTagInteractor(gateways.NewStatusTagGateway(dbconn), presenters.NewStatusTagPresenter())
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, resp := si.UpdateStatusTag(context.Background(), tc.input)
			if status != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, status)
			}
			require.Equal(t, resp, tc.want)
		})
	}
}
