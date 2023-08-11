package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/stretchr/testify/require"
)

// インテグレートテスト
func createHackathonStatusTagTest(t *testing.T, hackathon repository.Hackathon) repository.HackathonStatusTag {
	arg := repository.CreateHackathonStatusTagsParams{
		HackathonID: hackathon.HackathonID,
		StatusID:    int32(1),
	}

	statusTag, err := testQueries.CreateHackathonStatusTags(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, statusTag)

	require.Equal(t, arg.HackathonID, statusTag.HackathonID)
	require.Equal(t, arg.StatusID, statusTag.StatusID)

	return statusTag
}

func TestCreateHackathonStatusTag(t *testing.T) {
	hackathons := createHackathonTest(t)
	createHackathonStatusTagTest(t, hackathons)
}

func TestGetStatusTags(t *testing.T) {
	hackathons := createHackathonTest(t)

	n := 2
	for i := 0; i < n; i++ {
		createHackathonStatusTagTest(t, hackathons)
	}

	statusTags, err := testQueries.ListHackathonStatusTagsByID(context.Background(), hackathons.HackathonID)
	require.NoError(t, err)
	require.NotEmpty(t, statusTags)
	require.Len(t, statusTags, n)

	for _, statusTag := range statusTags {
		require.NotEmpty(t, statusTag)
		require.Equal(t, hackathons.HackathonID, statusTag.HackathonID)

	}
}
