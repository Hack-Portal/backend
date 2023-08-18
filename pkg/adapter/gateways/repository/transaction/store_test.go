package transaction

import (
	fb "firebase.google.com/go"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

func NewFakeStore(app *fb.App) *SQLStore {
	return &SQLStore{
		Queries: repository.NewFake(),
		App:     app,
	}
}
