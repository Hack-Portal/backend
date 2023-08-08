package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func createPastWorkTagsTest(t *testing.T, pastWork repository.PastWork) repository.PastWorkTag {
	tags := listTechTagTest(t)
	techs := util.Random(len(tags) - 1)
	arg := repository.CreatePastWorkTagsParams{
		Opus:      pastWork.Opus,
		TechTagID: tags[techs].TechTagID,
	}
	pastWorkTags, err := testQueries.CreatePastWorkTags(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pastWorkTags)

	require.Equal(t, arg.Opus, pastWorkTags.Opus)
	require.Equal(t, arg.TechTagID, pastWorkTags.TechTagID)

	return pastWorkTags
}

func TestCreatePastWorkTag(t *testing.T) {
	pastWork := createPastWorksTest(t)
	createPastWorkTagsTest(t, pastWork)
}

func TestListPastWorkTagsByOpus(t *testing.T) {
	n := 5
	pastWorks := createPastWorksTest(t)
	for i := 0; i < n; i++ {
		createPastWorkTagsTest(t, pastWorks)
	}
	pastWorkTags, err := testQueries.ListPastWorkTagsByOpus(context.Background(), pastWorks.Opus)
	require.NoError(t, err)
	require.NotEmpty(t, pastWorkTags)
	require.Len(t, pastWorkTags, n)
}

func TestDeletePastWorkTags(t *testing.T) {
	n := 5
	pastWorks := createPastWorksTest(t)
	for i := 0; i < n; i++ {
		createPastWorkTagsTest(t, pastWorks)
	}
	err := testQueries.DeletePastWorkTagsByOpus(context.Background(), pastWorks.Opus)
	require.NoError(t, err)
	pastWorkTags, err := testQueries.ListPastWorkTagsByOpus(context.Background(), pastWorks.Opus)
	require.NoError(t, err)
	require.Empty(t, pastWorkTags)
}
