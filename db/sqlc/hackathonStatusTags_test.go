package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// インテグレートテスト
func createHackathonStatusTagTest(t *testing.T, hackathon Hackathons) HackathonStatusTags {
	arg := CreateHackathonStatusTagParams{
		HackathonID: hackathon.HackathonID,
		StatusID:    int32(1),
	}

	statusTag, err := testQueries.CreateHackathonStatusTag(context.Background(), arg)
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

	statusTags, err := testQueries.GetHackathonStatusTagsByHackathonID(context.Background(), hackathons.HackathonID)
	require.NoError(t, err)
	require.NotEmpty(t, statusTags)
	require.Len(t, statusTags, n)

	for _, statusTag := range statusTags {
		require.NotEmpty(t, statusTag)
		require.Equal(t, hackathons.HackathonID, statusTag.HackathonID)

	}
}
