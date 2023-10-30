package gateways

import (
	"context"
	"temp/src/entities"
	"temp/src/entities/param"
	"temp/src/entities/request"
	"temp/src/usecase"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"gorm.io/gorm"
)

type hackathonGateway struct {
	db  *gorm.DB
	app *firebase.App
}

func NewHackathonGateway(db *gorm.DB, app *firebase.App) usecase.HackathonPort {
	return &hackathonGateway{
		db:  db,
		app: app,
	}
}

func (g *hackathonGateway) Create(arg param.CreateHackathon) error {
	return g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(arg.Hackathon).Error; err != nil {
			return err
		}
		// こっちの方が早い
		if err := createHackathonStatusTags(tx, arg.Hackathon.HackathonID, arg.StatusTags); err != nil {
			return err
		}

		return nil
	})
}

func createHackathonStatusTags(db *gorm.DB, hackathonID string, statusTags []int) error {
	if len(statusTags) == 0 {
		return nil
	}

	query := `INSERT INTO hackathon_status_tags (hackathon_id, status_id) VALUES `
	params := []interface{}{}

	for _, tag := range statusTags {
		query += `(?, ?),`
		params = append(params, hackathonID, tag)
	}

	// 末尾のカンマを削除
	query = query[:len(query)-1]
	query += ` RETURNING hackathon_id, status_id`

	if err := db.Raw(query, params...).Error; err != nil {
		return err
	}
	return nil
}

// 取得
func (g *hackathonGateway) Get(request.ListRequest) ([]*entities.Hackathon, error) { return nil, nil }

// 更新
func (g *hackathonGateway) UpdateByID(*entities.Hackathon) (*entities.Hackathon, error) {
	return nil, nil
}

// 削除
func (g *hackathonGateway) DeleteByID(int32) error { return nil }

// 承認
func (g *hackathonGateway) Approve(int32) error { return nil }

// 画像登録
func (g *hackathonGateway) UplaodImage(hackathonID string, image []byte) (filpath string, err error) {
	ctx := context.Background()
	bucket, err := g.selectDefaultBucket(ctx)
	if err != nil {
		return "", err
	}

	obj := bucket.Object(hackathonID + ".jpg")
	w := obj.NewWriter(ctx)

	if _, err := w.Write(image); err != nil {
		return "", err
	}

	if err := w.Close(); err != nil {
		return "", err
	}

	filpath, err = bucket.SignedURL(obj.ObjectName(), &storage.SignedURLOptions{
		Expires: time.Now().AddDate(100, 0, 0),
		Method:  "GET",
	})

	if err != nil {
		return "", err
	}

	return
}

// 画像削除
func (g *hackathonGateway) DeleteImage(string) error { return nil }

func (g *hackathonGateway) selectDefaultBucket(ctx context.Context) (*storage.BucketHandle, error) {
	fbstorage, err := g.app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	return fbstorage.DefaultBucket()
}
