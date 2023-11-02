package usecases

import (
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/utils"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/outputport"
)

// ここではUsecase Interactorを実装する

type HackathonInteractor struct {
	HackathonOutput     outputport.HackathonOutputPort
	HackathonRepository dai.HackathonRepository
	Firebase            dai.FirebaseRepository
}

func NewHackathonInterface(output outputport.HackathonOutputPort, repository dai.HackathonRepository) inputport.HackathonInputPort {
	return &HackathonInteractor{
		HackathonOutput:     output,
		HackathonRepository: repository,
	}
}

func (hi *HackathonInteractor) Create(arg input.HackathonCreate, icon []byte) {
	hackathonID := utils.NewUUID()

	image, err := hi.Firebase.UploadFile(hackathonID, icon)
	if err != nil {
		hi.HackathonOutput.Create(err)
		return
	}

	param := deformationHackathonCreate(hackathonID, arg, image)
	param.Statuses, err = utils.StrToIntArr(arg.StatusTags)
	if err != nil {
		hi.HackathonOutput.Create(err)
		return
	}

	if err = hi.HackathonRepository.Create(param); err != nil {
		hi.HackathonOutput.Create(err)
		return
	}

	hi.HackathonOutput.Create(nil)
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
