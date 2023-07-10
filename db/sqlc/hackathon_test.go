package db

import (
	"context"
	"testing"
	"time"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createHackathonTest(t *testing.T) Hackathons {

	arg := CreateHackathonParams{
		Name:        util.RandomString(8),
		Icon:        []byte(util.RandomString(8)),
		Description: util.RandomString(8),
		Link:        util.RandomString(8),
		// 時間を適当に生成すればいい
		// 今から10時間後の時間を返す
		// ex: time.Now().Add(time.Duration(time.Duration(10).Hours()))
		Expired:   time.Now().Add(time.Duration(time.Duration(100).Hours())),
		StartDate: time.Now().Add(time.Duration(time.Duration(200).Hours())),
		Term:      int32(util.Random(100)),
	}

	hackathon, err := testQueries.CreateHackathon(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hackathon)

	require.Equal(t, arg.Name, hackathon.Name)
	require.Equal(t, arg.Icon, hackathon.Icon)
	require.Equal(t, arg.Description, hackathon.Description)
	require.Equal(t, arg.Link, hackathon.Link)
	require.Equal(t, arg.Term, hackathon.Term)

	require.NotZero(t, hackathon.HackathonID)

	return hackathon
}

func TestCreateHackathon(t *testing.T) {
	createHackathonTest(t)
}
