package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createPastWorksTest(t *testing.T) PastWorks {
	arg := CreatePastWorksParams{
		Name:            util.RandomString(8),
		ThumbnailImage:  []byte(util.RandomString(8)),
		ExplanatoryText: util.RandomString(32),
	}
	past_work, err := testQueries.CreatePastWorks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, past_work)

	require.Equal(t, arg.Name, past_work.Name)
	require.Equal(t, arg.ThumbnailImage, past_work.ThumbnailImage)
	require.Equal(t, arg.ExplanatoryText, past_work.ExplanatoryText)

	require.NotZero(t, past_work.Opus)
	return past_work
}
func TestCreatePastWorks(t *testing.T) {
	createPastWorksTest(t)
}

func TestGetPastWorks(t *testing.T) {
	pastWork1 := createPastWorksTest(t)
	pastWork2, err := testQueries.GetPastWorksByOpus(context.Background(), pastWork1.Opus)
	require.NoError(t, err)
	require.NotEmpty(t, pastWork2)

	require.Equal(t, pastWork1.Opus, pastWork2.Opus)
	require.Equal(t, pastWork1.Name, pastWork2.Name)
	require.Equal(t, pastWork1.ThumbnailImage, pastWork2.ThumbnailImage)
	require.Equal(t, pastWork1.ExplanatoryText, pastWork2.ExplanatoryText)
}

func TestListPastWorks(t *testing.T) {
	n := 5

	for i := 0; i < n; i++ {
		createPastWorksTest(t)
	}

	arg := ListPastWorksParams{
		Limit:  int32(n),
		Offset: 0,
	}

	pastWorks, err := testQueries.ListPastWorks(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, pastWorks)
	require.Len(t, pastWorks, n)

	for _, pastWork := range pastWorks {
		require.NotEmpty(t, pastWork)
	}
}
