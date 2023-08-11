package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func createPastWorkFrameworksTest(t *testing.T, pastWork repository.PastWork) repository.PastWorkFramework {
	frameworks := listFrameworkTest(t)
	framework := util.Random(len(frameworks) - 1)
	arg := repository.CreatePastWorkFrameworksParams{
		Opus:        pastWork.Opus,
		FrameworkID: frameworks[framework].FrameworkID,
	}
	pastWorkFrameworks, err := testQueries.CreatePastWorkFrameworks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pastWorkFrameworks)

	require.Equal(t, arg.Opus, pastWorkFrameworks.Opus)
	require.Equal(t, arg.FrameworkID, pastWorkFrameworks.FrameworkID)

	return pastWorkFrameworks
}

func TestCreatePastWorkFrameworks(t *testing.T) {
	pastWork := createPastWorksTest(t)
	createPastWorkFrameworksTest(t, pastWork)
}

func TestListPastWorkFrameworksByOpus(t *testing.T) {
	n := 5
	pastWorks := createPastWorksTest(t)
	for i := 0; i < n; i++ {
		createPastWorkFrameworksTest(t, pastWorks)
	}
	pastWorkFrameworks, err := testQueries.ListPastWorkFrameworksByOpus(context.Background(), pastWorks.Opus)
	require.NoError(t, err)
	require.NotEmpty(t, pastWorkFrameworks)
	require.Len(t, pastWorkFrameworks, n)
}

func TestDeletePastWorkFrameworks(t *testing.T) {
	n := 5
	pastWorks := createPastWorksTest(t)
	for i := 0; i < n; i++ {
		createPastWorkFrameworksTest(t, pastWorks)
	}
	err := testQueries.DeletePastWorkFrameworksByOpus(context.Background(), pastWorks.Opus)
	require.NoError(t, err)
	pastWorkFrameworks, err := testQueries.ListPastWorkFrameworksByOpus(context.Background(), pastWorks.Opus)
	require.NoError(t, err)
	require.Empty(t, pastWorkFrameworks)
}
