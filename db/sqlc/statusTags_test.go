package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var statusData = []StatusTags{
	{1, "オンライン"},
}

func TestListStatusTags(t *testing.T) {
	statusTags, err := testQueries.ListStatusTags(context.Background(), int32(1))
	require.NoError(t, err)
	require.NotEmpty(t, statusTags)

	require.Equal(t, statusTags, statusData)
}
