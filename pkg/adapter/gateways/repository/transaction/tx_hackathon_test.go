package transaction

import (
	"context"
	"testing"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
	"github.com/stretchr/testify/require"
)

func randomHachathon(t *testing.T) (params.CreateHackathon, repository.Hackathon) {
	arg := params.CreateHackathon{
		Hackathon: repository.CreateHackathonsParams{
			Name:        util.RandomString(10),
			Icon:        dbutil.ToSqlNullString(util.RandomString(10)),
			Description: util.RandomString(10),
			Link:        util.RandomString(10),
			Expired:     time.Now().Add(time.Hour * 24),
			StartDate:   time.Now().Add(time.Hour * 48),
			Term:        10,
		},
		StatusTags: util.RandomSelection(3, 2),
	}

	hackathon, err := store.CreateHackathonTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hackathon)

	return arg, hackathon
}

func TestCreateHackathonTx(t *testing.T) {
	arg, hackathon := randomHachathon(t)

	tags, err := store.ListHackathonStatusTagsByID(context.Background(), hackathon.HackathonID)
	require.NoError(t, err)

	require.NotEmpty(t, hackathon.HackathonID)
	require.Equal(t, arg.Hackathon.Name, hackathon.Name)
	require.Equal(t, arg.Hackathon.Icon, hackathon.Icon)
	require.Equal(t, arg.Hackathon.Description, hackathon.Description)
	require.Equal(t, arg.Hackathon.Link, hackathon.Link)
	require.NotEmpty(t, hackathon.Expired)
	require.NotEmpty(t, hackathon.StartDate)
	require.Equal(t, arg.Hackathon.Term, hackathon.Term)
	require.Len(t, tags, len(arg.StatusTags))
}
