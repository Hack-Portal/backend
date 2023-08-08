package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func listFrameworkTest(t *testing.T) []repository.Framework {
	frameworks, err := testQueries.ListFrameworks(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, frameworks)
	// require.Len(t, frameworks, n)
	return frameworks
}

func TestListFrameworks(t *testing.T) {
	listFrameworkTest(t)
}

func TestGetFrameworks(t *testing.T) {
	frameworks := listFrameworkTest(t)
	randomId := util.Random(len(frameworks) - 1)

	framework, err := testQueries.GetFrameworksByID(context.Background(), frameworks[randomId].FrameworkID)
	require.NoError(t, err)
	require.NotEmpty(t, framework)

	require.Equal(t, frameworks[randomId].FrameworkID, framework.FrameworkID)
	require.Equal(t, frameworks[randomId].Framework, framework.Framework)
	require.Equal(t, frameworks[randomId].TechTagID, framework.TechTagID)
}
