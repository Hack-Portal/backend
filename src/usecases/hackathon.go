package usecases

import (
	"temp/pkg/utils"
	"temp/src/datastructs/entities"
	"temp/src/datastructs/input"
	"temp/src/datastructs/params"
	"temp/src/usecases/dai"
	"temp/src/usecases/inputport"
	"temp/src/usecases/outputport"
	"time"
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
	image, err := hi.Firebase.UploadFile(icon)
	if err != nil {
		hi.HackathonOutput.Create(err)
		return
	}
	param := deformationHackathonCreate(arg, image)
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

func deformationHackathonCreate(arg input.HackathonCreate, icon string) params.HackathonCreate {
	return params.HackathonCreate{
		Hackathon: entities.Hackathon{
			HackathonID: utils.NewUUID(),
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
