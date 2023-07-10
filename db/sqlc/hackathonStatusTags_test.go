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
	statusTag1 := createHackathonStatusTagTest(t)

	statusTag2, err := testQueries.GetStatusTags(context.Background(), int32(1))
	require.NoError(t, err)
	require.NotEmpty(t, statusTag2)

	require.Equal(t, statusTag1.HackathonID, statusTag2.HackathonID)
	require.Equal(t, statusTag1.StatusID, statusTag2.StatusID)
}
