package repository_test

import (
	"context"
	"testing"
	"time"

	domain "github.com/hackhack-Geek-vol6/backend/pkg/domain"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func createHackathonTest(t *testing.T) Hackathons {

	arg := domain.CreateHackathonRequestBody{
		Name:        util.RandomString(8),
		Description: util.RandomString(8),
		Link:        util.RandomString(8),
		// 時間を適当に生成すればいい
		// 今から10時間後の時間を返す
		// ex: time.Now().Add(time.Duration(time.Duration(10).Hours()))
		Expired:   time.Now().Add(time.Hour * 100),
		StartDate: time.Now().Add(time.Hour * 200),
		Term:      int32(util.Random(100)),
	}

	hackathon, err := testQueries.CreateHackathon(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hackathon)

	require.Equal(t, arg.Name, hackathon.Name)
	require.Equal(t, arg.Description, hackathon.Description)
	require.Equal(t, arg.Link, hackathon.Link)
	require.Equal(t, arg.Term, hackathon.Term)

	require.NotZero(t, hackathon.HackathonID)

	return hackathon
}

func TestCreateHackathon(t *testing.T) {
	createHackathonTest(t)
}

func TestGetHackathon(t *testing.T) {
	hackathon1 := createHackathonTest(t)

	hackathon2, err := testQueries.GetHackathonByID(context.Background(), hackathon1.HackathonID)
	require.NoError(t, err)
	require.NotEmpty(t, hackathon2)

	require.Equal(t, hackathon1.HackathonID, hackathon2.HackathonID)
	require.Equal(t, hackathon1.Name, hackathon2.Name)
	require.Equal(t, hackathon1.Icon, hackathon2.Icon)
	require.Equal(t, hackathon1.Description, hackathon2.Description)
	require.Equal(t, hackathon1.Link, hackathon2.Link)
	require.Equal(t, hackathon1.Expired, hackathon2.Expired)
	require.Equal(t, hackathon1.StartDate, hackathon2.StartDate)
	require.Equal(t, hackathon1.Term, hackathon2.Term)
}
func TestListHackathon(t *testing.T) {
	n := 5

	for i := 0; i < n; i++ {
		createHackathonTest(t)
	}
	arg := domain.ListHackathonsParams{
		Expired: time.Now(),
		Limit:   int32(n),
		Offset:  0,
	}

	hackathons, err := testQueries.ListHackathons(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, hackathons)
	require.Len(t, hackathons, n)

	for _, hackathon := range hackathons {
		require.NotEmpty(t, hackathon)
	}
}
