package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
)

func NewFake() *Queries {
	return &Queries{}
}

type fakeQuerier struct {
	// 1
	statusTag map[int32]StatusTag
	locate    map[int32]Locate
	techTag   map[int32]TechTag
	role      map[int32]Role
	hackathon map[int32]Hackathon
	// 2
	framework map[int32]Framework
	account   map[string]Account
	// 3
	room               map[string]Room
	rateEntity         map[int32]RateEntity
	follow             map[int32]Follow
	pastWork           map[int32]PastWork
	accountTag         map[int32]AccountTag
	accountFramework   map[int32]AccountFramework
	hackathonStatusTag map[int32]HackathonStatusTag
	// 4
	accountPastWork   map[int32]AccountPastWork
	like              map[int32]Like
	pastWorkFramework map[int32]PastWorkFramework
	pastWorkTag       map[int32]PastWorkTag
	roomsAccount      map[int32]RoomsAccount
}

func NewFakeDB() Querier {
	return &fakeQuerier{}
}

func (fq *fakeQuerier) CreateLocates(ctx context.Context, name string) (Locate, error) {
	locate := Locate{
		LocateID: int32(len(fq.locate) + 1),
		Name:     name,
	}

	if len(locate.Name) == 0 {
		err := errors.New(fmt.Sprintf(`null value in column "%s" violates not-null constraint`, "name"))
		return Locate{}, err
	}

	fq.locate[int32(len(fq.locate)+1)] = locate

	return fq.locate[int32(len(fq.locate)+1)], nil
}

func (fq *fakeQuerier) GetLocatesByID(ctx context.Context, locateID int32) (Locate, error) {
	locate, ok := fq.locate[locateID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, locateID))
		return Locate{}, err
	}
	return locate, nil
}

func (fq *fakeQuerier) ListLocates(ctx context.Context) ([]Locate, error) {
	locates := []Locate{}
	for _, locate := range fq.locate {
		locates = append(locates, locate)
	}
	return locates, nil
}

func (fq *fakeQuerier) CreateTechTags(ctx context.Context, language string) (TechTag, error) {
	techTag := TechTag{
		TechTagID: int32(len(fq.techTag)) + 1,
		Language:  language,
	}
	if len(techTag.Language) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "language"))
		return TechTag{}, err
	}

	fq.techTag[int32(len(fq.techTag))+1] = techTag
	return fq.techTag[int32(len(fq.techTag))+1], nil
}

func (fq *fakeQuerier) GetTechTagsByID(ctx context.Context, techTagID int32) (TechTag, error) {
	techTag, ok := fq.techTag[techTagID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, techTagID))
		return TechTag{}, err
	}
	return techTag, nil
}

func (fq *fakeQuerier) ListTechTags(ctx context.Context) ([]TechTag, error) {
	techTags := []TechTag{}

	for _, techTag := range fq.techTag {
		techTags = append(techTags, techTag)
	}

	return techTags, nil
}

func (fq *fakeQuerier) UpdateTechTagsByID(ctx context.Context, arg UpdateTechTagsByIDParams) (TechTag, error) {
	if _, ok := fq.techTag[arg.TechTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.TechTagID))
		return TechTag{}, err
	}
	techTag := TechTag{
		TechTagID: arg.TechTagID,
		Language:  arg.Language,
	}
	fq.techTag[arg.TechTagID] = techTag
	return fq.techTag[arg.TechTagID], nil
}

func (fq *fakeQuerier) DeleteTechTagsByID(ctx context.Context, techTagID int32) error {
	if _, ok := fq.techTag[techTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, techTagID))
		return err
	}
	delete(fq.techTag, techTagID)

	return nil
}

func (fq *fakeQuerier) CreateRoles(ctx context.Context, role string) (Role, error) {
	r := Role{
		RoleID: int32(len(fq.role)) + 1,
		Role:   role,
	}
	if len(r.Role) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "role"))
		return Role{}, err
	}
	fq.role[int32(len(fq.role))+1] = r
	return fq.role[int32(len(fq.role))+1], nil
}

func (fq *fakeQuerier) GetRolesByID(ctx context.Context, roleID int32) (Role, error) {
	role, ok := fq.role[roleID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, roleID))
		return Role{}, err
	}
	return role, nil
}

func (fq *fakeQuerier) ListRoles(ctx context.Context) ([]Role, error) {
	roles := []Role{}

	for _, role := range fq.role {
		roles = append(roles, role)
	}

	return roles, nil
}

func (fq *fakeQuerier) CreateStatusTags(ctx context.Context, status string) (StatusTag, error) {
	statusTag := StatusTag{
		StatusID: int32(len(fq.statusTag)) + 1,
		Status:   status,
	}
	if len(statusTag.Status) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "status"))
		return StatusTag{}, err
	}

	fq.statusTag[int32(len(fq.statusTag))+1] = statusTag
	return fq.statusTag[int32(len(fq.statusTag))+1], nil
}

func (fq *fakeQuerier) GetStatusTagsByTag(ctx context.Context, statusID int32) (StatusTag, error) {
	statuTag, ok := fq.statusTag[statusID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, statusID))
		return StatusTag{}, err
	}
	return statuTag, nil
}

func (fq *fakeQuerier) ListStatusTags(ctx context.Context) ([]StatusTag, error) {
	statusTags := []StatusTag{}

	for _, statusTag := range fq.statusTag {
		statusTags = append(statusTags, statusTag)
	}

	return statusTags, nil
}

func (fq *fakeQuerier) DeleteStatusTagsByStatusID(ctx context.Context, statusID int32) error {
	if _, ok := fq.statusTag[statusID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, statusID))
		return err
	}
	delete(fq.statusTag, statusID)

	return nil
}

func (fq *fakeQuerier) CreateHackathons(ctx context.Context, arg CreateHackathonsParams) (Hackathon, error) {
	if len(arg.Name) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "name"))
		return Hackathon{}, err
	}
	if len(arg.Link) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "link"))
		return Hackathon{}, err
	}
	hackathon := Hackathon{
		HackathonID: int32(len(fq.hackathon)) + 1,
		Name:        arg.Name,
		Icon:        arg.Icon,
		Description: arg.Description,
		Link:        arg.Link,
		Expired:     arg.Expired,
		StartDate:   arg.StartDate,
		Term:        arg.Term,
	}
	fq.hackathon[int32(len(fq.hackathon))+1] = hackathon
	return fq.hackathon[int32(len(fq.hackathon))+1], nil
}

func (fq *fakeQuerier) GetHackathonByID(ctx context.Context, hackathonID int32) (Hackathon, error) {
	hackathon, ok := fq.hackathon[hackathonID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, hackathonID))
		return Hackathon{}, err
	}
	return hackathon, nil
}

func (fq *fakeQuerier) ListHackathons(ctx context.Context, arg ListHackathonsParams) ([]Hackathon, error) {
	var count int32
	hackathons := []Hackathon{}

	for _, hackathon := range fq.hackathon {
		if arg.Offset*arg.Limit-arg.Limit > count {
			if hackathon.Expired.Sub(arg.Expired) > 0 {
				hackathons = append(hackathons, hackathon)
			}
		}

		count++
		if count >= arg.Limit {
			break
		}
	}

	return hackathons, nil
}

func (fq *fakeQuerier) DeleteHackathonByID(ctx context.Context, hackathonID int32) error {
	if _, ok := fq.hackathon[hackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, hackathonID))
		return err
	}
	delete(fq.hackathon, hackathonID)
	return nil
}

func (fq *fakeQuerier) CreateFrameworks(ctx context.Context, arg CreateFrameworksParams) (Framework, error) {
	if _, ok := fq.techTag[arg.TechTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.TechTagID))
		return Framework{}, err
	}
	if len(arg.Framework) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "framework"))
		return Framework{}, err
	}
	framework := Framework{
		FrameworkID: int32(len(fq.framework)) + 1,
		TechTagID:   arg.TechTagID,
		Framework:   arg.Framework,
	}
	fq.framework[int32(len(fq.framework))+1] = framework
	return fq.framework[int32(len(fq.framework))+1], nil
}

func (fq *fakeQuerier) GetFrameworksByID(ctx context.Context, frameworkID int32) (Framework, error) {
	framework, ok := fq.framework[frameworkID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, frameworkID))
		return Framework{}, err
	}
	return framework, nil
}

func (fq *fakeQuerier) ListFrameworks(ctx context.Context) ([]Framework, error) {
	frameworks := []Framework{}

	for _, framework := range fq.framework {
		frameworks = append(frameworks, framework)
	}

	return frameworks, nil
}

func (fq *fakeQuerier) UpdateFrameworksByID(ctx context.Context, arg UpdateFrameworksByIDParams) (Framework, error) {
	if _, ok := fq.framework[arg.FrameworkID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.FrameworkID))
		return Framework{}, err
	}
	if _, ok := fq.techTag[arg.TechTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.TechTagID))
		return Framework{}, err
	}

	fq.framework[arg.FrameworkID] = Framework{
		FrameworkID: arg.FrameworkID,
		TechTagID:   arg.TechTagID,
		Framework:   arg.Framework,
	}

	return fq.framework[arg.FrameworkID], nil
}
func (fq *fakeQuerier) DeleteFrameworksByID(ctx context.Context, frameworkID int32) error {
	if _, ok := fq.framework[frameworkID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, frameworkID))
		return err
	}
	delete(fq.framework, frameworkID)
	return nil
}

func (fq *fakeQuerier) CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error) {
	if len(arg.AccountID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return Account{}, err
	}
	if len(arg.Email) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "user_id"))
		return Account{}, err
	}
	if len(arg.Username) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "username"))
		return Account{}, err
	}
	if arg.LocateID == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "locate_id"))
		return Account{}, err
	}

	for _, account := range fq.account {
		if account.AccountID == arg.AccountID {
			err := errors.New(fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s" `, arg.AccountID))
			return Account{}, err
		}
	}

	for _, account := range fq.account {
		if account.Email == arg.Email {
			err := errors.New(fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s" `, arg.Email))
			return Account{}, err
		}
	}

	if _, ok := fq.locate[arg.LocateID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.LocateID))
		return Account{}, err
	}

	account := Account{
		AccountID:       arg.AccountID,
		Email:           arg.Email,
		Username:        arg.Username,
		Icon:            arg.Icon,
		ExplanatoryText: arg.ExplanatoryText,
		LocateID:        arg.LocateID,
		Rate:            arg.Rate,
		ShowLocate:      arg.ShowLocate,
		ShowRate:        arg.ShowRate,
		Character:       arg.Character,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDelete:        false,
	}

	fq.account[arg.AccountID] = account
	return fq.account[arg.AccountID], nil
}

func (fq *fakeQuerier) GetAccountsByID(ctx context.Context, accountID string) (Account, error) {
	account, ok := fq.account[accountID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, accountID))
		return Account{}, err
	}
	return account, nil
}

func (fq *fakeQuerier) GetAccountsByEmail(ctx context.Context, email string) (Account, error) {
	for _, account := range fq.account {
		if account.Email == account.Email {
			return account, nil
		}
	}

	err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, email))
	return Account{}, err
}

func (fq *fakeQuerier) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	var count int32
	accounts := []Account{}

	for _, account := range fq.account {
		if arg.Offset*arg.Limit-arg.Limit > count {
			accounts = append(accounts, account)
		}
		count++
		if count >= arg.Limit {
			break
		}
	}

	return accounts, nil
}

func (fq *fakeQuerier) UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) (Account, error) {
	account, ok := fq.account[arg.AccountID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return Account{}, err
	}

	if _, ok := fq.locate[arg.LocateID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.LocateID))
		return Account{}, err
	}

	newAccount := Account{
		AccountID:       account.AccountID,
		Email:           account.Email,
		Username:        arg.Username,
		Icon:            arg.Icon,
		ExplanatoryText: arg.ExplanatoryText,
		LocateID:        arg.LocateID,
		Rate:            arg.Rate,
		Character:       arg.Character,
		ShowLocate:      arg.ShowLocate,
		ShowRate:        arg.ShowRate,
		CreateAt:        account.CreateAt,
		UpdateAt:        time.Now(),
		IsDelete:        false,
	}

	fq.account[arg.AccountID] = newAccount

	return fq.account[arg.AccountID], nil
}

func (fq *fakeQuerier) UpdateRateByID(ctx context.Context, arg UpdateRateByIDParams) (Account, error) {
	account, ok := fq.account[arg.AccountID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return Account{}, err
	}

	newAccount := Account{
		AccountID:       account.AccountID,
		Email:           account.Email,
		Username:        account.Username,
		Icon:            account.Icon,
		ExplanatoryText: account.ExplanatoryText,
		LocateID:        account.LocateID,
		Rate:            arg.Rate,
		Character:       account.Character,
		ShowLocate:      account.ShowLocate,
		ShowRate:        account.ShowRate,
		CreateAt:        account.CreateAt,
		UpdateAt:        time.Now(),
		IsDelete:        false,
	}
	fq.account[arg.AccountID] = newAccount

	return fq.account[arg.AccountID], nil
}

func (fq *fakeQuerier) DeleteAccounts(ctx context.Context, accountID string) (Account, error) {
	if len(accountID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return Account{}, err
	}
	account := fq.account[accountID]
	account.IsDelete = true
	fq.account[accountID] = account
	return fq.account[accountID], nil
}

func (fq *fakeQuerier) CreateRooms(ctx context.Context, arg CreateRoomsParams) (Room, error) {
	if len(arg.RoomID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "room_id"))
		return Room{}, err
	}
	if len(arg.Title) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "title"))
		return Room{}, err
	}
	if len(arg.Description) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "description"))
		return Room{}, err
	}

	if _, ok := fq.hackathon[arg.HackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.HackathonID))
		return Room{}, err
	}

	room := Room{
		RoomID:      arg.RoomID,
		HackathonID: arg.HackathonID,
		Title:       arg.Title,
		Description: arg.Description,
		MemberLimit: arg.MemberLimit,
		IncludeRate: arg.IncludeRate,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
		IsDelete:    false,
	}
	fq.room[arg.RoomID] = room
	return fq.room[arg.RoomID], nil
}

func (fq *fakeQuerier) GetRoomsByID(ctx context.Context, roomID string) (Room, error) {
	room, ok := fq.room[roomID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, roomID))
		return Room{}, err
	}
	return room, nil
}

func (fq *fakeQuerier) ListRooms(ctx context.Context, arg ListRoomsParams) ([]Room, error) {
	var count int32
	rooms := []Room{}

	for _, room := range fq.room {
		if arg.Offset*arg.Limit-arg.Limit > count {
			rooms = append(rooms, room)
		}
		count++
		if count >= arg.Limit {
			break
		}
	}

	return rooms, nil
}

func (fq *fakeQuerier) UpdateRoomsByID(ctx context.Context, arg UpdateRoomsByIDParams) (Room, error) {
	if len(arg.RoomID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "room_id"))
		return Room{}, err
	}
	if len(arg.Title) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "title"))
		return Room{}, err
	}
	if len(arg.Description) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "description"))
		return Room{}, err
	}

	room, ok := fq.room[arg.RoomID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.HackathonID))
		return Room{}, err
	}

	if _, ok := fq.hackathon[arg.HackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.HackathonID))
		return Room{}, err
	}

	newRoom := Room{
		RoomID:      arg.RoomID,
		HackathonID: arg.HackathonID,
		Title:       arg.Title,
		Description: arg.Description,
		MemberLimit: arg.MemberLimit,
		IncludeRate: room.IncludeRate,
		CreateAt:    room.CreateAt,
		UpdateAt:    time.Now(),
		IsDelete:    false,
	}

	fq.room[arg.RoomID] = newRoom
	return fq.room[arg.RoomID], nil
}

func (fq *fakeQuerier) DeleteRoomsByID(ctx context.Context, roomID string) (Room, error) {
	if len(roomID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "room_id"))
		return Room{}, err
	}
	room := fq.room[roomID]
	room.IsDelete = true
	fq.room[roomID] = room
	return fq.room[roomID], nil
}

func (fq *fakeQuerier) CreateRateEntities(ctx context.Context, arg CreateRateEntitiesParams) (RateEntity, error) {
	if len(arg.AccountID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return RateEntity{}, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return RateEntity{}, err
	}
	rate := RateEntity{
		AccountID: arg.AccountID,
		Rate:      arg.Rate,
		CreateAt:  time.Now(),
	}
	fq.rateEntity[int32(len(fq.rateEntity))+1] = rate
	return fq.rateEntity[int32(len(fq.rateEntity))+1], nil
}

func (fq *fakeQuerier) ListRateEntities(ctx context.Context, arg ListRateEntitiesParams) ([]RateEntity, error) {
	if len(arg.AccountID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return nil, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return nil, err
	}

	var count int32
	rates := []RateEntity{}
	for _, rate := range fq.rateEntity {
		if arg.Offset*arg.Limit-arg.Limit > count {
			rates = append(rates, rate)
		}
		count++
		if count >= arg.Limit {
			break
		}
	}
	return rates, nil
}

func (fq *fakeQuerier) CreateFollows(ctx context.Context, arg CreateFollowsParams) (Follow, error) {
	if _, ok := fq.account[arg.ToAccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.ToAccountID))
		return Follow{}, err
	}

	if _, ok := fq.account[arg.FromAccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.FromAccountID))
		return Follow{}, err
	}

	follow := Follow{
		ToAccountID:   arg.ToAccountID,
		FromAccountID: arg.FromAccountID,
		CreateAt:      time.Now(),
	}

	fq.follow[int32(len(fq.follow))+1] = follow
	return fq.follow[int32(len(fq.follow))+1], nil
}

func (fq *fakeQuerier) DeleteFollows(ctx context.Context, arg DeleteFollowsParams) error {
	if _, ok := fq.account[arg.ToAccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.ToAccountID))
		return err
	}

	if _, ok := fq.account[arg.FromAccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.FromAccountID))
		return err
	}
	for i, follow := range fq.follow {
		if follow.FromAccountID == arg.FromAccountID && follow.ToAccountID == arg.ToAccountID {
			delete(fq.follow, i)
			return nil
		}
	}

	return errors.New(fmt.Sprintf(`ERROR: no rows in result set `))
}

func (fq *fakeQuerier) ListFollowsByFromUserID(ctx context.Context, fromAccountID string) ([]Follow, error) {
	follows := []Follow{}
	for i, follow := range fq.follow {
		if follow.FromAccountID == fromAccountID {
			follows = append(follows, fq.follow[i])
		}
	}
	return follows, nil
}

func (fq *fakeQuerier) ListFollowsByToUserID(ctx context.Context, toAccountID string) ([]Follow, error) {
	follows := []Follow{}
	for i, follow := range fq.follow {
		if follow.ToAccountID == toAccountID {
			follows = append(follows, fq.follow[i])
		}
	}
	return follows, nil
}

func (fq *fakeQuerier) CreatePastWorks(ctx context.Context, arg CreatePastWorksParams) (PastWork, error) {
	if len(arg.Name) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return PastWork{}, err
	}

	if len(arg.ThumbnailImage) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return PastWork{}, err
	}

	if len(arg.ExplanatoryText) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return PastWork{}, err
	}

	pastWork := PastWork{
		Opus:            int32(len(fq.pastWork)) + 1,
		Name:            arg.Name,
		ThumbnailImage:  arg.ThumbnailImage,
		ExplanatoryText: arg.ExplanatoryText,
		AwardDataID:     arg.AwardDataID,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDelete:        false,
	}
	fq.pastWork[int32(len(fq.pastWork))+1] = pastWork
	return fq.pastWork[int32(len(fq.pastWork))+1], nil
}

func (fq *fakeQuerier) GetPastWorksByOpus(ctx context.Context, opus int32) (PastWork, error) {
	pastwork, ok := fq.pastWork[opus]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return PastWork{}, err
	}

	return pastwork, nil
}

func (fq *fakeQuerier) ListPastWorks(ctx context.Context, arg ListPastWorksParams) ([]ListPastWorksRow, error) {
	pastWorks := []ListPastWorksRow{}

	var count int32
	for _, pastWork := range fq.pastWork {
		if arg.Offset*arg.Limit-arg.Limit > count {
			pastWorks = append(pastWorks, ListPastWorksRow{
				Opus:            pastWork.Opus,
				Name:            pastWork.Name,
				ExplanatoryText: pastWork.ExplanatoryText,
			})
		}
		count++
		if count >= arg.Limit {
			break
		}
	}
	return pastWorks, nil
}

func (fq *fakeQuerier) UpdatePastWorksByID(ctx context.Context, arg UpdatePastWorksByIDParams) (PastWork, error) {
	if len(arg.Name) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "room_id"))
		return PastWork{}, err
	}
	if len(arg.ThumbnailImage) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "title"))
		return PastWork{}, err
	}
	if len(arg.ExplanatoryText) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "description"))
		return PastWork{}, err
	}

	pastWork, ok := fq.pastWork[arg.Opus]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return PastWork{}, err
	}

	newPastWork := PastWork{
		Opus:            arg.Opus,
		Name:            arg.Name,
		ThumbnailImage:  arg.ThumbnailImage,
		ExplanatoryText: arg.ExplanatoryText,
		AwardDataID:     arg.AwardDataID,
		CreateAt:        pastWork.CreateAt,
		UpdateAt:        time.Now(),
		IsDelete:        false,
	}

	fq.pastWork[arg.Opus] = newPastWork
	return fq.pastWork[arg.Opus], nil
}

func (fq *fakeQuerier) DeletePastWorksByID(ctx context.Context, arg DeletePastWorksByIDParams) (PastWork, error) {
	pastWork, ok := fq.pastWork[arg.Opus]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return PastWork{}, err
	}
	pastWork.IsDelete = true

	fq.pastWork[arg.Opus] = pastWork
	return fq.pastWork[arg.Opus], nil
}

func (fq *fakeQuerier) CreateAccountTags(ctx context.Context, arg CreateAccountTagsParams) (AccountTag, error) {
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return AccountTag{}, err
	}
	if _, ok := fq.techTag[arg.TechTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.TechTagID))
		return AccountTag{}, err
	}

	accountTag := AccountTag{
		AccountID: arg.AccountID,
		TechTagID: arg.TechTagID,
	}

	fq.accountTag[int32(len(fq.accountTag))+1] = accountTag
	return fq.accountTag[int32(len(fq.accountTag))+1], nil
}

func (fq *fakeQuerier) ListAccountTagsByUserID(ctx context.Context, accountID string) ([]ListAccountTagsByUserIDRow, error) {
	if _, ok := fq.account[accountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, accountID))
		return nil, err
	}
	accountTags := []ListAccountTagsByUserIDRow{}

	for _, accountTag := range fq.accountTag {
		if accountTag.AccountID == accountID {
			tag := fq.techTag[accountTag.TechTagID]
			accountTags = append(accountTags, ListAccountTagsByUserIDRow{
				TechTagID: dbutil.ToSqlNullInt32(tag.TechTagID),
				Language:  dbutil.ToSqlNullString(tag.Language),
			})
		}
	}

	return accountTags, nil
}

func (fq *fakeQuerier) DeleteAccountTagsByUserID(ctx context.Context, accountID string) error {
	if _, ok := fq.account[accountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, accountID))
		return err
	}

	for i, accountTag := range fq.accountTag {
		if accountTag.AccountID == accountID {
			delete(fq.accountTag, i)
		}
	}

	return nil
}

func (fq *fakeQuerier) CreateAccountFrameworks(ctx context.Context, arg CreateAccountFrameworksParams) (AccountFramework, error) {
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return AccountFramework{}, err
	}
	if _, ok := fq.framework[arg.FrameworkID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.FrameworkID))
		return AccountFramework{}, err
	}

	accountFramework := AccountFramework{
		AccountID:   arg.AccountID,
		FrameworkID: arg.FrameworkID,
	}

	fq.accountFramework[int32(len(fq.accountFramework))+1] = accountFramework
	return fq.accountFramework[int32(len(fq.accountFramework))+1], nil
}

func (fq *fakeQuerier) ListAccountFrameworksByUserID(ctx context.Context, accountID string) ([]ListAccountFrameworksByUserIDRow, error) {
	if _, ok := fq.account[accountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, accountID))
		return nil, err
	}
	accountFramework := []ListAccountFrameworksByUserIDRow{}

	for _, accountTag := range fq.accountFramework {
		if accountTag.AccountID == accountID {
			tag := fq.framework[accountTag.FrameworkID]
			accountFramework = append(accountFramework, ListAccountFrameworksByUserIDRow{
				TechTagID:   dbutil.ToSqlNullInt32(tag.TechTagID),
				FrameworkID: dbutil.ToSqlNullInt32(tag.FrameworkID),
				Framework:   dbutil.ToSqlNullString(tag.Framework),
			})
		}
	}

	return accountFramework, nil
}

func (fq *fakeQuerier) DeleteAccountFrameworkByUserID(ctx context.Context, accountID string) error {
	if _, ok := fq.account[accountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, accountID))
		return err
	}

	for i, accountFramework := range fq.accountFramework {
		if accountFramework.AccountID == accountID {
			delete(fq.accountTag, i)
		}
	}

	return nil
}

func (fq *fakeQuerier) CreateHackathonStatusTags(ctx context.Context, arg CreateHackathonStatusTagsParams) (HackathonStatusTag, error) {
	if _, ok := fq.hackathon[arg.HackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.HackathonID))
		return HackathonStatusTag{}, err
	}
	if _, ok := fq.statusTag[arg.StatusID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.StatusID))
		return HackathonStatusTag{}, err
	}

	hackathonStatusTag := HackathonStatusTag{
		HackathonID: arg.HackathonID,
		StatusID:    arg.StatusID,
	}

	fq.hackathonStatusTag[int32(len(fq.hackathonStatusTag))+1] = hackathonStatusTag
	return fq.hackathonStatusTag[int32(len(fq.hackathonStatusTag))+1], nil
}

func (fq *fakeQuerier) ListHackathonStatusTagsByID(ctx context.Context, hackathonID int32) ([]HackathonStatusTag, error) {
	if _, ok := fq.hackathon[hackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, hackathonID))
		return nil, err
	}

	hackathonTags := []HackathonStatusTag{}

	for _, hackathonTag := range fq.hackathonStatusTag {
		if hackathonTag.HackathonID == hackathonID {
			hackathonTags = append(hackathonTags, HackathonStatusTag{
				HackathonID: hackathonTag.HackathonID,
				StatusID:    hackathonTag.StatusID,
			})
		}
	}
	return hackathonTags, nil
}

func (fq *fakeQuerier) DeleteHackathonStatusTagsByID(ctx context.Context, hackathonID int32) error {
	if _, ok := fq.hackathon[hackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, hackathonID))
		return err
	}

	for i, hackathonTag := range fq.hackathonStatusTag {
		if hackathonTag.HackathonID == hackathonID {
			delete(fq.hackathonStatusTag, i)
		}
	}

	return nil
}

func (fq *fakeQuerier) CreateAccountPastWorks(ctx context.Context, arg CreateAccountPastWorksParams) (AccountPastWork, error) {
	if _, ok := fq.pastWork[arg.Opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return AccountPastWork{}, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return AccountPastWork{}, err
	}

	accountPastWork := AccountPastWork{
		Opus:      arg.Opus,
		AccountID: arg.AccountID,
	}

	fq.accountPastWork[int32(len(fq.accountPastWork))+1] = accountPastWork
	return fq.accountPastWork[int32(len(fq.accountPastWork))+1], nil
}

func (fq *fakeQuerier) ListAccountPastWorksByOpus(ctx context.Context, opus int32) ([]AccountPastWork, error) {
	if _, ok := fq.pastWork[opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return nil, err
	}

	accountPastWorks := []AccountPastWork{}

	for _, accountPastWork := range fq.accountPastWork {
		if accountPastWork.Opus == opus {
			accountPastWorks = append(accountPastWorks, AccountPastWork{
				Opus:      accountPastWork.Opus,
				AccountID: accountPastWork.AccountID,
			})
		}
	}
	return accountPastWorks, nil
}

func (fq *fakeQuerier) DeleteAccountPastWorksByOpus(ctx context.Context, opus int32) error {
	if _, ok := fq.pastWork[opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return err
	}

	for i, accountPastWork := range fq.accountPastWork {
		if accountPastWork.Opus == opus {
			delete(fq.accountPastWork, i)
		}
	}
	return nil
}

func (fq *fakeQuerier) CreatePastWorkFrameworks(ctx context.Context, arg CreatePastWorkFrameworksParams) (PastWorkFramework, error) {
	if _, ok := fq.pastWorkFramework[arg.Opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return PastWorkFramework{}, err
	}
	if _, ok := fq.framework[arg.FrameworkID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.FrameworkID))
		return PastWorkFramework{}, err
	}

	pastWorkFramework := PastWorkFramework{
		Opus:        arg.Opus,
		FrameworkID: arg.FrameworkID,
	}

	fq.pastWorkFramework[int32(len(fq.pastWorkFramework))+1] = pastWorkFramework
	return fq.pastWorkFramework[int32(len(fq.pastWorkFramework))+1], nil
}

func (fq *fakeQuerier) ListPastWorkFrameworksByOpus(ctx context.Context, opus int32) ([]PastWorkFramework, error) {
	if _, ok := fq.pastWork[opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return nil, err
	}

	pastWorkFrameworks := []PastWorkFramework{}

	for _, pastWorkFramework := range fq.pastWorkFramework {
		if pastWorkFramework.Opus == opus {

			pastWorkFrameworks = append(pastWorkFrameworks, PastWorkFramework{
				Opus:        pastWorkFramework.Opus,
				FrameworkID: pastWorkFramework.FrameworkID,
			})
		}
	}
	return pastWorkFrameworks, nil
}

func (fq *fakeQuerier) DeletePastWorkFrameworksByOpus(ctx context.Context, opus int32) error {
	if _, ok := fq.pastWork[opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return err
	}

	for i, pastWorkFramework := range fq.pastWorkFramework {
		if pastWorkFramework.Opus == opus {
			delete(fq.pastWorkFramework, i)
		}
	}
	return nil
}

func (fq *fakeQuerier) CreatePastWorkTags(ctx context.Context, arg CreatePastWorkTagsParams) (PastWorkTag, error) {
	if _, ok := fq.pastWorkFramework[arg.Opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return PastWorkTag{}, err
	}
	if _, ok := fq.techTag[arg.TechTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.TechTagID))
		return PastWorkTag{}, err
	}

	pastWorkTag := PastWorkTag{
		Opus:      arg.Opus,
		TechTagID: arg.TechTagID,
	}

	fq.pastWorkTag[int32(len(fq.pastWorkTag))+1] = pastWorkTag
	return fq.pastWorkTag[int32(len(fq.pastWorkTag))+1], nil
}

func (fq *fakeQuerier) ListPastWorkTagsByOpus(ctx context.Context, opus int32) ([]PastWorkTag, error) {
	if _, ok := fq.pastWork[opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return nil, err
	}

	pastWorkTags := []PastWorkTag{}

	for _, pastWorkTag := range fq.pastWorkTag {
		if pastWorkTag.Opus == opus {

			pastWorkTags = append(pastWorkTags, PastWorkTag{
				Opus:      pastWorkTag.Opus,
				TechTagID: pastWorkTag.TechTagID,
			})
		}
	}
	return pastWorkTags, nil
}

func (fq *fakeQuerier) DeletePastWorkTagsByOpus(ctx context.Context, opus int32) error {
	if _, ok := fq.pastWork[opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return err
	}

	for i, pastWorkTag := range fq.pastWorkTag {
		if pastWorkTag.Opus == opus {
			delete(fq.pastWorkTag, i)
		}
	}
	return nil
}

func (fq *fakeQuerier) CreateRoomsAccounts(ctx context.Context, arg CreateRoomsAccountsParams) (RoomsAccount, error) {
	if _, ok := fq.room[arg.RoomID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.RoomID))
		return RoomsAccount{}, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return RoomsAccount{}, err
	}

	roomsAccount := RoomsAccount{
		AccountID: arg.AccountID,
		RoomID:    arg.RoomID,
		Role:      arg.Role,
		IsOwner:   arg.IsOwner,
		CreateAt:  time.Now(),
	}

	fq.roomsAccount[int32(len(fq.roomsAccount))+1] = roomsAccount
	return fq.roomsAccount[int32(len(fq.roomsAccount))+1], nil
}

func (fq *fakeQuerier) GetRoomsAccountsByID(ctx context.Context, roomID string) ([]GetRoomsAccountsByIDRow, error) {
	if _, ok := fq.room[roomID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, roomID))
		return nil, err
	}

	roomsAccounts := []GetRoomsAccountsByIDRow{}

	for _, roomsAccount := range fq.roomsAccount {
		if roomsAccount.RoomID == roomID {
			account := fq.account[roomsAccount.AccountID]

			roomsAccounts = append(roomsAccounts, GetRoomsAccountsByIDRow{
				AccountID: dbutil.ToSqlNullString(roomsAccount.AccountID),
				IsOwner:   roomsAccount.IsOwner,
				Icon:      account.Icon,
				Role:      roomsAccount.Role,
			})
		}
	}
	return roomsAccounts, nil
}

func (fq *fakeQuerier) DeleteRoomsAccountsByID(ctx context.Context, arg DeleteRoomsAccountsByIDParams) error {
	if _, ok := fq.room[arg.RoomID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.RoomID))
		return err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return err
	}

	for i, roomsAccount := range fq.roomsAccount {
		if roomsAccount.RoomID == arg.RoomID && roomsAccount.AccountID == arg.AccountID {
			delete(fq.pastWorkTag, i)
		}
	}
	return nil
}

func (fq *fakeQuerier) CreateLikes(ctx context.Context, arg CreateLikesParams) (Like, error) {
	if _, ok := fq.pastWork[arg.Opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return Like{}, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return Like{}, err
	}

	like := Like{
		Opus:      arg.Opus,
		AccountID: arg.AccountID,
		CreateAt:  time.Now(),
		IsDelete:  false,
	}
	fq.like[int32(len(fq.like))+1] = like
	return fq.like[int32(len(fq.like))+1], nil
}

func (fq *fakeQuerier) GetLikeStatusByID(ctx context.Context, arg GetLikeStatusByIDParams) (Like, error) {
	if _, ok := fq.pastWork[arg.Opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return Like{}, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return Like{}, err
	}

	for i, like := range fq.like {
		if like.Opus == arg.Opus && like.AccountID == arg.AccountID {
			return fq.like[i], nil
		}
	}

	return Like{}, errors.New(fmt.Sprintf(`ERROR: no rows in result set `))
}

func (fq *fakeQuerier) GetListCountByOpus(ctx context.Context, opus int32) (int64, error) {
	var count int64
	if _, ok := fq.pastWork[opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return 0, err
	}

	for _, like := range fq.like {
		if like.Opus == opus {
			count++
		}
	}
	return count, nil
}

func (fq *fakeQuerier) ListLikesByID(ctx context.Context, accountID string) ([]Like, error) {
	if _, ok := fq.account[accountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, accountID))
		return nil, err
	}

	likes := []Like{}

	for i, like := range fq.like {
		if like.AccountID == accountID {
			likes = append(likes, fq.like[i])
		}
	}
	return likes, nil
}

func (fq *fakeQuerier) DeleteLikesByID(ctx context.Context, arg DeleteLikesByIDParams) (Like, error) {
	if _, ok := fq.pastWork[arg.Opus]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return Like{}, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return Like{}, err
	}

	for i, like := range fq.like {
		if like.Opus == arg.Opus && like.AccountID == arg.AccountID {
			like := fq.like[i]
			like.IsDelete = true
			fq.like[i] = like
			return fq.like[i], nil
		}
	}
	return Like{}, errors.New(fmt.Sprintf(`ERROR: no rows in result set `))
}
