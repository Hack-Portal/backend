package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createHackathonStatusTagTest(t *testing.T) HackathonStatusTags {
	hackathon := createHackathonTest(t)

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
	createHackathonStatusTagTest(t)
}

func TestGetStatusTags(t *testing.T) {
	Hackathons := createHackathonTest(t)
	n := 5
	for i := 0; i < n; i++ {
		createHackathonStatusTagTest(t)
	}

	statusTag, err := testQueries.GetStatusTags(context.Background(), Hackathons.HackathonID)
	require.NoError(t, err)

	for _, statusTag := range statusTag {
		require.NotEmpty(t, statusTag)
		require.Equal(t, Hackathons.HackathonID, statusTag.HackathonID)
		require.Len(t, statusTag.StatusID, n)
	}
}
