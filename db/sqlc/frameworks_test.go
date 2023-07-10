package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func listFrameworkTest(t *testing.T) []Frameworks {
	n := 5
	frameworks, err := testQueries.ListFrameworks(context.Background(), int32(n))
	require.NoError(t, err)
	require.NotEmpty(t, frameworks)
	require.Len(t, frameworks, n)
	return frameworks
}

func TestListFrameworks(t *testing.T) {
	listFrameworkTest(t)
}

func TestGetFrameworks(t *testing.T) {
	frameworks := listFrameworkTest(t)
	randomId := util.Random(len(frameworks) - 1)

	framework, err := testQueries.GetFrameworks(context.Background(), frameworks[randomId].FrameworkID)
	require.NoError(t, err)
	require.NotEmpty(t, framework)

	require.Equal(t, frameworks[randomId].FrameworkID, framework.FrameworkID)
	require.Equal(t, frameworks[randomId].Framework, framework.Framework)
	require.Equal(t, frameworks[randomId].TechTagID, framework.TechTagID)
}
