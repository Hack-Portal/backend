package usecase

import (
	"temp/pkg/utils"
	"temp/src/entities"
	"temp/src/entities/param"
	"temp/src/entities/request"
	"temp/src/usecase/input"
	"temp/src/usecase/output"
	"time"

	"github.com/google/uuid"
)

type HackathonPort interface {
	// 登録
	Create(param.CreateHackathon) error
	// 取得
	Get(request.ListRequest) ([]*entities.Hackathon, error)
	// 更新
	UpdateByID(*entities.Hackathon) (*entities.Hackathon, error)
	// 削除
	DeleteByID(int32) error

	// 承認
	Approve(int32) error

	// 画像登録
	UplaodImage(string, []byte) (string, error)
	// 画像削除
	DeleteImage(string) error
}

type Hackathon struct {
	outputPort    output.HackathonOutputPort
	hackathonPort HackathonPort
}

func NewHackathon(op output.HackathonOutputPort, hp HackathonPort) input.HackathonInputPort {
	return &Hackathon{
		outputPort:    op,
		hackathonPort: hp,
	}
}

func (h *Hackathon) Create(arg request.CreateHackathon, image []byte) error {
	icon, err := h.hackathonPort.UplaodImage(arg.Name, image)
	if err != nil {
		return h.outputPort.RenderError(err)
	}

	statusTags, err := utils.StrToIntArr(arg.StatusTags)
	if err != nil {
		return err
	}

	param := transformHackathonParam(arg, icon)
	param.StatusTags = statusTags

	if err := h.hackathonPort.Create(param); err != nil {
		return h.outputPort.RenderError(err)
	}

	return h.outputPort.RenderCreate()
}

func (h *Hackathon) Get(arg request.ListRequest) ([]*entities.Hackathon, error) {
	hackathons, err := h.hackathonPort.Get(arg)
	if err != nil {
		return nil, h.outputPort.RenderError(err)
	}

	return hackathons, h.outputPort.RenderGet(hackathons)
}

func (h *Hackathon) UpdateByID(hackathon *entities.Hackathon) error { return nil }
func (h *Hackathon) DeleteByID(hackathonID int32) error             { return nil }
func (h *Hackathon) Approve(hackathonID int32) error                { return nil }

func transformHackathonParam(h request.CreateHackathon, icon string) param.CreateHackathon {
	return param.CreateHackathon{
		Hackathon: &entities.Hackathon{
			HackathonID: uuid.New().String(),
			Name:        h.Name,
			Description: h.Description,
			Icon:        icon,
			StartDate:   h.StartDate,
			Link:        h.Link,
			Expired:     h.Expired,
			Term:        h.Term,
			CreatedAt:   time.Now(),
			IsDelete:    false,
		},
	}
}
