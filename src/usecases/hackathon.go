package usecases

import (
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/utils"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/cerror"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/output"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/ports"
)

// ここではUsecase Interactorを実装する

type HackathonInteractor struct {
	Output              ports.HackathonOutputBoundary
	HackathonRepository dai.HackathonRepository
	Firebase            dai.FirebaseRepository
}

func NewHackathonInteractor(
	output ports.HackathonOutputBoundary,
	repository dai.HackathonRepository,
	firebase dai.FirebaseRepository,
) ports.HackathonInputBoundary {
	return &HackathonInteractor{
		Output:              output,
		HackathonRepository: repository,
		Firebase:            firebase,
	}
}

func (hi *HackathonInteractor) Create(arg input.HackathonCreate, icon []byte) (int, *output.CreateHackathon) {
	hackathonID := utils.NewUUID()

	if icon == nil {
		return hi.Output.Create(cerror.ImageNotFound)
	}

	image, err := hi.Firebase.UploadFile(hackathonID, icon)
	if err != nil {
		return hi.Output.Create(err)
	}

	param := deformationHackathonCreate(hackathonID, arg, image)
	param.Statuses, err = utils.StrToIntArr(arg.StatusTags)
	if err != nil {
		return hi.Output.Create(err)
	}

	if err = hi.HackathonRepository.Create(param); err != nil {
		return hi.Output.Create(err)
	}

	return hi.Output.Create(nil)
}

func (hi *HackathonInteractor) Read()   {}
func (hi *HackathonInteractor) Update() {}
func (hi *HackathonInteractor) Delete() {}

func deformationHackathonCreate(hackathonID string, arg input.HackathonCreate, icon string) params.HackathonCreate {
	return params.HackathonCreate{
		Hackathon: entities.Hackathon{
			HackathonID: hackathonID,
			Name:        arg.Name,
			Icon:        icon,
			StartDate:   arg.StartDate,
			Link:        arg.Link,
			Expired:     arg.Expired,
			Term:        arg.Term,
			CreatedAt:   time.Now(),
			IsDelete:    false,
		},
	}
}
