package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/stretchr/testify/require"
)

func listTechTagTest(t *testing.T) []repository.TechTag {
	techTags, err := testQueries.ListTechTags(context.Background())
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

	techTag, err := testQueries.GetTechTagsByID(context.Background(), techTags[n].TechTagID)
	require.NoError(t, err)
	require.NotEmpty(t, techTag)

	require.Equal(t, techTags[n].TechTagID, techTag.TechTagID)
	require.Equal(t, techTags[n].Language, techTag.Language)
}
