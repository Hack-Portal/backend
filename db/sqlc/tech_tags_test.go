package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func listTechTagTest(t *testing.T) []TechTags {
	techTags, err := testQueries.ListTechTag(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, techTags)
	return techTags
}

func TestListTechTag(t *testing.T) {
	listTechTagTest(t)
}

func TestGetTechTag(t *testing.T) {
	n := 1
	techTags := listTechTagTest(t)

	techTag, err := testQueries.GetTechTag(context.Background(), techTags[n].TechTagID)
	require.NoError(t, err)
	require.NotEmpty(t, techTag)

	require.Equal(t, techTags[n].TechTagID, techTag.TechTagID)
	require.Equal(t, techTags[n].Language, techTag.Language)
}
