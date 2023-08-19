package transaction

import (
	"context"
	"database/sql"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func TestCreatePastWorkTx(t *testing.T) {
	store := NewStore(testDB, App)
	var accountID []string
	for i := 0; i < 3; i++ {
		_, account := randomAccount(t, store)
		accountID = append(accountID, account.AccountID)
	}

	arg := domain.CreatePastWorkParams{
		Name:            util.RandomString(10),
		ThumbnailImage:  util.RandomString(10),
		ExplanatoryText: util.RandomString(10),
		//TODO:AwardData追加APIが追加されたら変更する
		AwardDataID:        sql.NullInt32{Valid: false},
		PastWorkTags:       util.RandomSelection(14, 3),
		PastWorkFrameworks: util.RandomSelection(51, 3),
		AccountPastWorks:   accountID,
	}

	pastwork, err := store.CreatePastWorkTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pastwork)

	tag, err := store.ListPastWorkTagsByOpus(context.Background(), pastwork.Opus)
	require.NoError(t, err)

	framework, err := store.ListPastWorkFrameworksByOpus(context.Background(), pastwork.Opus)
	require.NoError(t, err)

	require.NotZero(t, pastwork.Opus)
	require.Equal(t, arg.Name, pastwork.Name)
	require.Equal(t, arg.ThumbnailImage, pastwork.ThumbnailImage)
	require.Equal(t, arg.ExplanatoryText, pastwork.ExplanatoryText)
	require.Equal(t, arg.AwardDataID, pastwork.AwardDataID)
	require.NotZero(t, pastwork.CreateAt)
	require.NotZero(t, pastwork.UpdateAt)
	require.Len(t, tag, len(arg.PastWorkTags))
	require.Len(t, framework, len(arg.PastWorkFrameworks))
}
