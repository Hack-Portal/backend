package interactors

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/google/uuid"
)

const (
	// HACKATHON_IMAGE_DIR はハッカソンの画像を保存するディレクトリ
	HACKATHON_IMAGE_DIR = "hackathon/"
)

type HackathonInteractor struct {
	Hackathon       dai.HackathonDai
	HackathonStatus dai.HackathonStatusDai
	FileStore       dai.FileStore
	HackathonOutput ports.HackathonOutputBoundary
}

func NewHackathonInteractor(hackathonDai dai.HackathonDai, HackathonStatus dai.HackathonStatusDai, filestoreDai dai.FileStore, hackathonOutput ports.HackathonOutputBoundary) ports.HackathonInputBoundary {
	return &HackathonInteractor{
		Hackathon:       hackathonDai,
		HackathonStatus: HackathonStatus,
		FileStore:       filestoreDai,
		HackathonOutput: hackathonOutput,
	}
}

func (hi *HackathonInteractor) CreateHackathon(ctx context.Context, in *ports.InputCreatehackathonData) (int, *response.CreateHackathon) {
	// 画像があるときは画像を保存してLinkを追加
	// 画像がないときは初期画像をLinkに追加
	var (
		imageLinks  string = config.Config.Server.DefaultHackathonImage
		hackathonID string = uuid.New().String()
	)

	if in.ImageFile != nil {
		src, err := in.ImageFile.Open()
		if err != nil {
			return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
				Error:    err,
				Response: nil,
			})
		}
		defer src.Close()

		data, err := io.ReadAll(src)
		if err != nil {
			return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
				Error:    err,
				Response: nil,
			})
		}

		// 画像を保存してLinkを追加
		links, err := hi.FileStore.UploadFile(ctx, data, fmt.Sprintf("%s%s.%s", HACKATHON_IMAGE_DIR, hackathonID, in.ImageFile.Filename))
		if err != nil {
			return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
				Error:    err,
				Response: nil,
			})
		}

		imageLinks = links
	}

	// ハッカソンを作成
	if err := hi.Hackathon.Create(ctx, &models.Hackathon{
		HackathonID: hackathonID,
		Name:        in.Name,
		Link:        in.Link,
		Expired:     in.Expired,
		StartDate:   in.StartDate,
		Term:        in.Term,
		Icon:        imageLinks,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}, in.Statuses); err != nil {
		return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
			Error:    err,
			Response: nil,
		})
	}

	// TODO:ハッカソンを取得？

	hackathon, status, err := hi.getHackathon(ctx, hackathonID)
	if err != nil {
		return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
			Error:    err,
			Response: nil,
		})
	}

	return hi.HackathonOutput.PresentCreateHackathon(ctx, &ports.OutputCreateHackathonData{
		Error: nil,
		Response: &response.CreateHackathon{
			HackathonID: hackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired.Format("2006-01-02"),
			StartDate:   hackathon.StartDate.Format("2006-01-02"),
			Term:        hackathon.Term,

			StatusTags: status,
		},
	})
}

func (hi *HackathonInteractor) GetHackathon(ctx context.Context, hackathonID string) (int, *response.GetHackathon) {
	if len(hackathonID) == 0 {
		return hi.HackathonOutput.PresentGetHackathon(ctx, &ports.OutputGetHackathonData{
			Error:    fmt.Errorf("invalid hackathon id"),
			Response: nil,
		})
	}

	hackathon, status, err := hi.getHackathon(ctx, hackathonID)
	if err != nil {
		return hi.HackathonOutput.PresentGetHackathon(ctx, &ports.OutputGetHackathonData{
			Error:    err,
			Response: nil,
		})
	}

	return hi.HackathonOutput.PresentGetHackathon(ctx, &ports.OutputGetHackathonData{
		Error: nil,
		Response: &response.GetHackathon{
			HackathonID: hackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired.Format("2006-01-02"),
			StartDate:   hackathon.StartDate.Format("2006-01-02"),
			Term:        hackathon.Term,

			StatusTags: status,
		},
	})
}

func (hi *HackathonInteractor) ListHackathon(ctx context.Context, pageID, pageSize int) (int, []*response.GetHackathon) {
	if pageID <= 0 {
		pageID = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	hackathons, err := hi.Hackathon.FindAll(ctx, pageSize, (pageID-1)*pageSize)
	if err != nil {
		return hi.HackathonOutput.PresentListHackathon(ctx, &ports.OutputListHackathonData{
			Error:    err,
			Response: nil,
		})
	}

	var parallelGetPresignedObjectURLInput []dai.ParallelGetPresignedObjectURLInput
	for _, hackathon := range hackathons {
		parallelGetPresignedObjectURLInput = append(parallelGetPresignedObjectURLInput, dai.ParallelGetPresignedObjectURLInput{
			HackathonID: hackathon.HackathonID,
			Key:         hackathon.Icon,
		})
	}
	icons, err := hi.FileStore.ParallelGetPresignedObjectURL(ctx, parallelGetPresignedObjectURLInput)
	if err != nil {
		return hi.HackathonOutput.PresentListHackathon(ctx, &ports.OutputListHackathonData{
			Error:    err,
			Response: nil,
		})
	}

	for _, hackathon := range hackathons {
		hackathon.Icon = icons[hackathon.HackathonID]
	}

	var hackathonIDs []string
	for _, hackathon := range hackathons {
		hackathonIDs = append(hackathonIDs, hackathon.HackathonID)
	}

	statuses, err := hi.HackathonStatus.FindAll(ctx, hackathonIDs)
	if err != nil {
		return hi.HackathonOutput.PresentListHackathon(ctx, &ports.OutputListHackathonData{
			Error:    err,
			Response: nil,
		})
	}

	var statusMap = make(map[string][]*response.StatusTag)
	for _, status := range statuses {
		statusMap[status.HackathonID] = append(statusMap[status.HackathonID], &response.StatusTag{
			ID:     status.StatusID,
			Status: status.Status,
		})
	}

	var responseHackathons []*response.GetHackathon
	for _, hackathon := range hackathons {
		responseHackathons = append(responseHackathons, &response.GetHackathon{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired.Format("2006-01-02"),
			StartDate:   hackathon.StartDate.Format("2006-01-02"),
			Term:        hackathon.Term,

			StatusTags: statusMap[hackathon.HackathonID],
		})
	}

	return hi.HackathonOutput.PresentListHackathon(ctx, &ports.OutputListHackathonData{
		Error:    nil,
		Response: responseHackathons,
	})
}

func (hi *HackathonInteractor) getHackathon(ctx context.Context, hackathonID string) (*models.Hackathon, []*response.StatusTag, error) {
	hackathon, err := hi.Hackathon.Find(ctx, hackathonID)
	if err != nil {
		return nil, nil, err
	}

	icon, err := hi.FileStore.GetPresignedObjectURL(ctx, hackathon.Icon)

	hackathon.Icon = icon

	statuses, err := hi.HackathonStatus.FindAll(ctx, []string{hackathonID})
	if err != nil {
		return nil, nil, err
	}

	var status []*response.StatusTag
	for _, s := range statuses {
		status = append(status, &response.StatusTag{
			ID:     s.StatusID,
			Status: s.Status,
		})
	}
	return hackathon, status, nil
}
