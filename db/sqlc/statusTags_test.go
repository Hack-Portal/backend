package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var statusData = []StatusTags{
	{1, "オンライン"},
	{2, "オフライン"},
	{3, "初心者歓迎"},
	{4, "急募"},
}

func TestListStatusTags(t *testing.T) {

	statusTags, err := testQueries.ListStatusTags(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, statusTags)

	require.Equal(t, statusTags, statusData)
}
