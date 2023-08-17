package fake

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

type fakeQuerier struct {
	// 1
	statusTag map[int32]repository.StatusTag
	locate    map[int32]repository.Locate
	techTag   map[int32]repository.TechTag
	role      map[int32]repository.Role
	user      map[string]repository.User
	hackathon map[int32]repository.Hackathon
	// 2
	framework map[int32]repository.Framework
	account   map[string]repository.Account
	// 3
	room               map[string]repository.Room
	rateEntity         map[int32]repository.RateEntity
	follow             map[int32]repository.Follow
	pastWork           map[int32]repository.PastWork
	accountTag         map[string]repository.AccountTag
	accountFramework   map[string]repository.AccountFramework
	hackathonStatusTag map[string]repository.HackathonStatusTag
	// 4
	accountPastWork   map[string]repository.AccountPastWork
	like              map[string]repository.Like
	pastWorkFramework map[string]repository.PastWorkFramework
	pastWorkTag       map[string]repository.PastWorkTag
	roomsAccount      map[string]repository.RoomsAccount
}

func (fq fakeQuerier) CreateLocates(ctx context.Context, name string) (repository.Locate, error) {
	locate := repository.Locate{
		LocateID: int32(len(fq.locate) + 1),
		Name:     name,
	}

	if len(locate.Name) == 0 {
		err := errors.New(fmt.Sprintf(`null value in column "%s" violates not-null constraint`, "name"))
		return repository.Locate{}, err
	}

	fq.locate[int32(len(fq.locate)+1)] = locate

	return fq.locate[int32(len(fq.locate)+1)], nil
}

func (fq fakeQuerier) GetLocatesByID(ctx context.Context, locateID int32) (repository.Locate, error) {
	locate, ok := fq.locate[locateID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, locateID))
		return repository.Locate{}, err
	}
	return locate, nil
}

func (fq fakeQuerier) ListLocates(ctx context.Context) ([]repository.Locate, error) {
	locates := []repository.Locate{}
	for _, locate := range fq.locate {
		locates = append(locates, locate)
	}
	return locates, nil
}

func (fq fakeQuerier) CreateTechTags(ctx context.Context, language string) (repository.TechTag, error) {
	techTag := repository.TechTag{
		TechTagID: int32(len(fq.techTag)) + 1,
		Language:  language,
	}
	if len(techTag.Language) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "language"))
		return repository.TechTag{}, err
	}

	fq.techTag[int32(len(fq.techTag))+1] = techTag
	return fq.techTag[int32(len(fq.techTag))+1], nil
}

func (fq fakeQuerier) GetTechTagsByID(ctx context.Context, techTagID int32) (repository.TechTag, error) {
	techTag, ok := fq.techTag[techTagID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, techTagID))
		return repository.TechTag{}, err
	}
	return techTag, nil
}

func (fq fakeQuerier) ListTechTags(ctx context.Context) ([]repository.TechTag, error) {
	techTags := []repository.TechTag{}

	for _, techTag := range fq.techTag {
		techTags = append(techTags, techTag)
	}

	return techTags, nil
}

func (fq fakeQuerier) UpdateTechTagsByID(ctx context.Context, arg repository.UpdateTechTagsByIDParams) (repository.TechTag, error) {
	if _, ok := fq.techTag[arg.TechTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.TechTagID))
		return repository.TechTag{}, err
	}
	techTag := repository.TechTag{
		TechTagID: arg.TechTagID,
		Language:  arg.Language,
	}
	fq.techTag[arg.TechTagID] = techTag
	return fq.techTag[arg.TechTagID], nil
}

func (fq fakeQuerier) DeleteTechTagsByID(ctx context.Context, techTagID int32) error {
	if _, ok := fq.techTag[techTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, techTagID))
		return err
	}
	delete(fq.techTag, techTagID)

	return nil
}

func (fq fakeQuerier) CreateRoles(ctx context.Context, role string) (repository.Role, error) {
	r := repository.Role{
		RoleID: int32(len(fq.role)) + 1,
		Role:   role,
	}
	if len(r.Role) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "role"))
		return repository.Role{}, err
	}
	fq.role[int32(len(fq.role))+1] = r
	return fq.role[int32(len(fq.role))+1], nil
}

func (fq fakeQuerier) GetRolesByID(ctx context.Context, roleID int32) (repository.Role, error) {
	role, ok := fq.role[roleID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, roleID))
		return repository.Role{}, err
	}
	return role, nil
}

func (fq fakeQuerier) ListRoles(ctx context.Context) ([]repository.Role, error) {
	roles := []repository.Role{}

	for _, role := range fq.role {
		roles = append(roles, role)
	}

	return roles, nil
}

func (fq fakeQuerier) CreateStatusTags(ctx context.Context, status string) (repository.StatusTag, error) {
	statusTag := repository.StatusTag{
		StatusID: int32(len(fq.statusTag)) + 1,
		Status:   status,
	}
	if len(statusTag.Status) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "status"))
		return repository.StatusTag{}, err
	}

	fq.statusTag[int32(len(fq.statusTag))+1] = statusTag
	return fq.statusTag[int32(len(fq.statusTag))+1], nil
}

func (fq fakeQuerier) GetStatusTagsByTag(ctx context.Context, statusID int32) (repository.StatusTag, error) {
	statuTag, ok := fq.statusTag[statusID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, statusID))
		return repository.StatusTag{}, err
	}
	return statuTag, nil
}

func (fq fakeQuerier) ListStatusTags(ctx context.Context) ([]repository.StatusTag, error) {
	statusTags := []repository.StatusTag{}

	for _, statusTag := range fq.statusTag {
		statusTags = append(statusTags, statusTag)
	}

	return statusTags, nil
}

func (fq fakeQuerier) DeleteStatusTagsByStatusID(ctx context.Context, statusID int32) error {
	if _, ok := fq.statusTag[statusID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, statusID))
		return err
	}
	delete(fq.statusTag, statusID)

	return nil
}

func (fq fakeQuerier) CreateUsers(ctx context.Context, arg repository.CreateUsersParams) (repository.User, error) {
	if len(arg.UserID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "user_id"))
		return repository.User{}, err
	}
	// Emailが空白でない時　重複がないかを確認する
	if arg.Email.Valid {
		for _, user := range fq.user {
			if user.Email.String == arg.Email.String {
				err := errors.New(fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s" `, arg.Email.String))
				return repository.User{}, err
			}
		}
	}

	for _, user := range fq.user {
		if user.UserID == arg.UserID {
			err := errors.New(fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s" `, arg.UserID))
			return repository.User{}, err
		}
	}

	user := repository.User{
		UserID:         arg.UserID,
		Email:          arg.Email,
		HashedPassword: arg.HashedPassword,
		CreateAt:       time.Now(),
		UpdateAt:       time.Now(),
		IsDelete:       false,
	}

	fq.user[arg.UserID] = user
	return fq.user[arg.UserID], nil
}

func (fq fakeQuerier) GetUsersByID(ctx context.Context, userID string) (repository.User, error) {
	user, ok := fq.user[userID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, userID))
		return repository.User{}, err
	}
	return user, nil
}

func (fq fakeQuerier) GetUsersByEmail(ctx context.Context, email sql.NullString) (repository.User, error) {
	for _, user := range fq.user {
		if user.Email.String == email.String {
			return user, nil
		}
	}

	err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, email.String))
	return repository.User{}, err
}

func (fq fakeQuerier) UpdateUsersByID(ctx context.Context, arg repository.UpdateUsersByIDParams) (repository.User, error) {
	if len(arg.UserID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "user_id"))
		return repository.User{}, err
	}

	user, ok := fq.user[arg.UserID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.UserID))
		return repository.User{}, err
	}

	if arg.Email.Valid {
		for _, user := range fq.user {
			if user.Email.String == arg.Email.String {
				err := errors.New(fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s" `, arg.Email.String))
				return repository.User{}, err
			}
		}
	}

	newUser := repository.User{
		UserID:         arg.UserID,
		Email:          arg.Email,
		HashedPassword: arg.HashedPassword,
		CreateAt:       user.CreateAt,
		UpdateAt:       time.Now(),
		IsDelete:       false,
	}

	fq.user[arg.UserID] = newUser

	return fq.user[arg.UserID], nil
}

func (fq fakeQuerier) DeleteUsersByID(ctx context.Context, arg repository.DeleteUsersByIDParams) error {
	if len(arg.UserID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "user_id"))
		return err
	}
	user := fq.user[arg.UserID]
	user.IsDelete = true
	fq.user[arg.UserID] = user
	return nil
}

func (fq fakeQuerier) CreateHackathons(ctx context.Context, arg repository.CreateHackathonsParams) (repository.Hackathon, error) {
	if len(arg.Name) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "name"))
		return repository.Hackathon{}, err
	}
	if len(arg.Link) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "link"))
		return repository.Hackathon{}, err
	}
	hackathon := repository.Hackathon{
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

func (fq fakeQuerier) GetHackathonByID(ctx context.Context, hackathonID int32) (repository.Hackathon, error) {
	hackathon, ok := fq.hackathon[hackathonID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, hackathonID))
		return repository.Hackathon{}, err
	}
	return hackathon, nil
}

func (fq fakeQuerier) ListHackathons(ctx context.Context, arg repository.ListHackathonsParams) ([]repository.Hackathon, error) {
	var count int32
	hackathons := []repository.Hackathon{}

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

func (fq fakeQuerier) DeleteHackathonByID(ctx context.Context, hackathonID int32) error {
	if _, ok := fq.hackathon[hackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, hackathonID))
		return err
	}
	delete(fq.hackathon, hackathonID)
	return nil
}

func (fq fakeQuerier) CreateFrameworks(ctx context.Context, arg repository.CreateFrameworksParams) (repository.Framework, error) {
	if _, ok := fq.techTag[arg.TechTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.TechTagID))
		return repository.Framework{}, err
	}
	if len(arg.Framework) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "framework"))
		return repository.Framework{}, err
	}
	framework := repository.Framework{
		FrameworkID: int32(len(fq.framework)) + 1,
		TechTagID:   arg.TechTagID,
		Framework:   arg.Framework,
	}
	fq.framework[int32(len(fq.framework))+1] = framework
	return fq.framework[int32(len(fq.framework))+1], nil
}

func (fq fakeQuerier) GetFrameworksByID(ctx context.Context, frameworkID int32) (repository.Framework, error) {
	framework, ok := fq.framework[frameworkID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, frameworkID))
		return repository.Framework{}, err
	}
	return framework, nil
}

func (fq fakeQuerier) ListFrameworks(ctx context.Context) ([]repository.Framework, error) {
	frameworks := []repository.Framework{}

	for _, framework := range fq.framework {
		frameworks = append(frameworks, framework)
	}

	return frameworks, nil
}

func (fq fakeQuerier) UpdateFrameworksByID(ctx context.Context, arg repository.UpdateFrameworksByIDParams) (repository.Framework, error) {
	if _, ok := fq.framework[arg.FrameworkID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.FrameworkID))
		return repository.Framework{}, err
	}
	if _, ok := fq.techTag[arg.TechTagID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.TechTagID))
		return repository.Framework{}, err
	}

	fq.framework[arg.FrameworkID] = repository.Framework{
		FrameworkID: arg.FrameworkID,
		TechTagID:   arg.TechTagID,
		Framework:   arg.Framework,
	}

	return fq.framework[arg.FrameworkID], nil
}
func (fq fakeQuerier) DeleteFrameworksByID(ctx context.Context, frameworkID int32) error {
	if _, ok := fq.framework[frameworkID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, frameworkID))
		return err
	}
	delete(fq.framework, frameworkID)
	return nil
}

func (fq fakeQuerier) CreateAccounts(ctx context.Context, arg repository.CreateAccountsParams) (repository.Account, error) {
	if len(arg.AccountID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return repository.Account{}, err
	}
	if len(arg.UserID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "user_id"))
		return repository.Account{}, err
	}
	if len(arg.Username) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "username"))
		return repository.Account{}, err
	}
	if arg.LocateID == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "locate_id"))
		return repository.Account{}, err
	}

	for _, account := range fq.account {
		if account.UserID == arg.UserID {
			err := errors.New(fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s" `, arg.UserID))
			return repository.Account{}, err
		}
	}

	if _, ok := fq.locate[arg.LocateID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.LocateID))
		return repository.Account{}, err
	}

	if _, ok := fq.user[arg.UserID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.UserID))
		return repository.Account{}, err
	}

	account := repository.Account{
		AccountID:       arg.AccountID,
		UserID:          arg.UserID,
		Username:        arg.Username,
		Icon:            arg.Icon,
		ExplanatoryText: arg.ExplanatoryText,
		LocateID:        arg.LocateID,
		Rate:            arg.Rate,
		ShowLocate:      arg.ShowLocate,
		ShowRate:        arg.ShowRate,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDelete:        false,
	}

	fq.account[arg.AccountID] = account
	return fq.account[arg.AccountID], nil
}

func (fq fakeQuerier) GetAccountsByID(ctx context.Context, accountID string) (repository.Account, error) {
	account, ok := fq.account[accountID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, accountID))
		return repository.Account{}, err
	}
	return account, nil
}

func (fq fakeQuerier) GetAccountsByEmail(ctx context.Context, email sql.NullString) (repository.Account, error) {
	user, err := fq.GetUsersByEmail(ctx, email)
	if err != nil {
		return repository.Account{}, err
	}

	for _, account := range fq.account {
		if user.UserID == account.UserID {
			return account, nil
		}
	}

	err = errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, email.String))
	return repository.Account{}, err
}

func (fq fakeQuerier) ListAccounts(ctx context.Context, arg repository.ListAccountsParams) ([]repository.Account, error) {
	var count int32
	accounts := []repository.Account{}

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

func (fq fakeQuerier) UpdateAccounts(ctx context.Context, arg repository.UpdateAccountsParams) (repository.Account, error) {
	account, ok := fq.account[arg.AccountID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return repository.Account{}, err
	}

	if _, ok := fq.locate[arg.LocateID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.LocateID))
		return repository.Account{}, err
	}

	newAccount := repository.Account{
		AccountID:       account.AccountID,
		UserID:          account.UserID,
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

func (fq fakeQuerier) UpdateRateByID(ctx context.Context, arg repository.UpdateRateByIDParams) (repository.Account, error) {
	account, ok := fq.account[arg.AccountID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return repository.Account{}, err
	}

	newAccount := repository.Account{
		AccountID:       account.AccountID,
		UserID:          account.UserID,
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

func (fq fakeQuerier) DeleteAccounts(ctx context.Context, accountID string) (repository.Account, error) {
	if len(accountID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return repository.Account{}, err
	}
	account := fq.account[accountID]
	account.IsDelete = true
	fq.account[accountID] = account
	return fq.account[accountID], nil
}

func (fq fakeQuerier) CreateRooms(ctx context.Context, arg repository.CreateRoomsParams) (repository.Room, error) {
	if len(arg.RoomID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "room_id"))
		return repository.Room{}, err
	}
	if len(arg.Title) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "title"))
		return repository.Room{}, err
	}
	if len(arg.Description) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "description"))
		return repository.Room{}, err
	}

	if _, ok := fq.hackathon[arg.HackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.HackathonID))
		return repository.Room{}, err
	}

	room := repository.Room{
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

func (fq fakeQuerier) GetRoomsByID(ctx context.Context, roomID string) (repository.Room, error) {
	room, ok := fq.room[roomID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, roomID))
		return repository.Room{}, err
	}
	return room, nil
}

func (fq fakeQuerier) ListRooms(ctx context.Context, arg repository.ListRoomsParams) ([]repository.Room, error) {
	var count int32
	rooms := []repository.Room{}

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

func (fq fakeQuerier) UpdateRoomsByID(ctx context.Context, arg repository.UpdateRoomsByIDParams) (repository.Room, error) {
	if len(arg.RoomID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "room_id"))
		return repository.Room{}, err
	}
	if len(arg.Title) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "title"))
		return repository.Room{}, err
	}
	if len(arg.Description) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "description"))
		return repository.Room{}, err
	}

	room, ok := fq.room[arg.RoomID]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.HackathonID))
		return repository.Room{}, err
	}

	if _, ok := fq.hackathon[arg.HackathonID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.HackathonID))
		return repository.Room{}, err
	}

	newRoom := repository.Room{
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

func (fq fakeQuerier) DeleteRoomsByID(ctx context.Context, roomID string) (repository.Room, error) {
	if len(roomID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "room_id"))
		return repository.Room{}, err
	}
	room := fq.room[roomID]
	room.IsDelete = true
	fq.room[roomID] = room
	return fq.room[roomID], nil
}

func (fq fakeQuerier) CreateRateEntities(ctx context.Context, arg repository.CreateRateEntitiesParams) (repository.RateEntity, error) {
	if len(arg.AccountID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return repository.RateEntity{}, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return repository.RateEntity{}, err
	}
	rate := repository.RateEntity{
		AccountID: arg.AccountID,
		Rate:      arg.Rate,
		CreateAt:  time.Now(),
	}
	fq.rateEntity[int32(len(fq.rateEntity))+1] = rate
	return fq.rateEntity[int32(len(fq.rateEntity))+1], nil
}

func (fq fakeQuerier) ListRateEntities(ctx context.Context, arg repository.ListRateEntitiesParams) ([]repository.RateEntity, error) {
	if len(arg.AccountID) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return nil, err
	}
	if _, ok := fq.account[arg.AccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.AccountID))
		return nil, err
	}

	var count int32
	rates := []repository.RateEntity{}
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

func (fq fakeQuerier) CreateFollows(ctx context.Context, arg repository.CreateFollowsParams) (repository.Follow, error) {
	if _, ok := fq.account[arg.ToAccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.ToAccountID))
		return repository.Follow{}, err
	}

	if _, ok := fq.account[arg.FromAccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.FromAccountID))
		return repository.Follow{}, err
	}

	follow := repository.Follow{
		ToAccountID:   arg.ToAccountID,
		FromAccountID: arg.FromAccountID,
		CreateAt:      time.Now(),
	}

	fq.follow[int32(len(fq.follow))+1] = follow
	return fq.follow[int32(len(fq.follow))+1], nil
}

func (fq fakeQuerier) DeleteFollows(ctx context.Context, arg repository.DeleteFollowsParams) error {
	if _, ok := fq.account[arg.ToAccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.ToAccountID))
		return err
	}

	if _, ok := fq.account[arg.FromAccountID]; !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%s" does not exist`, arg.FromAccountID))
		return err
	}
	for i, follow := range fq.follow {
		if follow.FromAccountID == arg.FromAccountID || follow.ToAccountID == arg.ToAccountID {
			delete(fq.follow, i)
			return nil
		}
	}

	return errors.New(fmt.Sprintf(`ERROR: no rows in result set `))
}

func (fq fakeQuerier) ListFollowsByFromUserID(ctx context.Context, fromAccountID string) ([]repository.Follow, error) {
	follows := []repository.Follow{}
	for i, follow := range fq.follow {
		if follow.FromAccountID == fromAccountID {
			follows = append(follows, fq.follow[i])
		}
	}
	return follows, nil
}

func (fq fakeQuerier) ListFollowsByToUserID(ctx context.Context, toAccountID string) ([]repository.Follow, error) {
	follows := []repository.Follow{}
	for i, follow := range fq.follow {
		if follow.ToAccountID == toAccountID {
			follows = append(follows, fq.follow[i])
		}
	}
	return follows, nil
}

func (fq fakeQuerier) CreatePastWorks(ctx context.Context, arg repository.CreatePastWorksParams) (repository.PastWork, error) {
	if len(arg.Name) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return repository.PastWork{}, err
	}

	if len(arg.ThumbnailImage) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return repository.PastWork{}, err
	}

	if len(arg.ExplanatoryText) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "account_id"))
		return repository.PastWork{}, err
	}

	pastWork := repository.PastWork{
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

func (fq fakeQuerier) GetPastWorksByOpus(ctx context.Context, opus int32) (repository.PastWork, error) {
	pastwork, ok := fq.pastWork[opus]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, opus))
		return repository.PastWork{}, err
	}

	return pastwork, nil
}

func (fq fakeQuerier) ListPastWorks(ctx context.Context, arg repository.ListPastWorksParams) ([]repository.ListPastWorksRow, error) {
	pastWorks := []repository.ListPastWorksRow{}

	var count int32
	for _, pastWork := range fq.pastWork {
		if arg.Offset*arg.Limit-arg.Limit > count {
			pastWorks = append(pastWorks, repository.ListPastWorksRow{
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

func (fq fakeQuerier) UpdatePastWorksByID(ctx context.Context, arg repository.UpdatePastWorksByIDParams) (repository.PastWork, error) {
	if len(arg.Name) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "room_id"))
		return repository.PastWork{}, err
	}
	if len(arg.ThumbnailImage) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "title"))
		return repository.PastWork{}, err
	}
	if len(arg.ExplanatoryText) == 0 {
		err := errors.New(fmt.Sprintf(`ERROR: null value in column "%s" violates not-null constraint`, "description"))
		return repository.PastWork{}, err
	}

	pastWork, ok := fq.pastWork[arg.Opus]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return repository.PastWork{}, err
	}

	newPastWork := repository.PastWork{
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

func (fq fakeQuerier) DeletePastWorksByID(ctx context.Context, arg repository.DeletePastWorksByIDParams) (repository.PastWork, error) {
	pastWork, ok := fq.pastWork[arg.Opus]
	if !ok {
		err := errors.New(fmt.Sprintf(`ERROR: column "%d" does not exist`, arg.Opus))
		return repository.PastWork{}, err
	}
	pastWork.IsDelete = false

	fq.pastWork[arg.Opus] = pastWork
	return fq.pastWork[arg.Opus], nil
}

func (fq fakeQuerier) CreateAccountTags(ctx context.Context, arg repository.CreateAccountTagsParams) (repository.AccountTag, error)
func (fq fakeQuerier) DeleteAccountTagsByUserID(ctx context.Context, accountID string) error
func (fq fakeQuerier) ListAccountTagsByUserID(ctx context.Context, accountID string) ([]repository.ListAccountTagsByUserIDRow, error)

func (fq fakeQuerier) CreateAccountFrameworks(ctx context.Context, arg repository.CreateAccountFrameworksParams) (repository.AccountFramework, error)
func (fq fakeQuerier) DeleteAccountFrameworkByUserID(ctx context.Context, accountID string) error
func (fq fakeQuerier) ListAccountFrameworksByUserID(ctx context.Context, accountID string) ([]repository.ListAccountFrameworksByUserIDRow, error)

func (fq fakeQuerier) CreateHackathonStatusTags(ctx context.Context, arg repository.CreateHackathonStatusTagsParams) (repository.HackathonStatusTag, error)
func (fq fakeQuerier) DeleteHackathonStatusTagsByID(ctx context.Context, hackathonID int32) error
func (fq fakeQuerier) ListHackathonStatusTagsByID(ctx context.Context, hackathonID int32) ([]repository.HackathonStatusTag, error)
func (fq fakeQuerier) GetStatusTagsByHackathonID(ctx context.Context, hackathonID int32) (repository.StatusTag, error)

func (fq fakeQuerier) CreateAccountPastWorks(ctx context.Context, arg repository.CreateAccountPastWorksParams) (repository.AccountPastWork, error)
func (fq fakeQuerier) DeleteAccountPastWorksByOpus(ctx context.Context, opus int32) error
func (fq fakeQuerier) ListAccountPastWorksByOpus(ctx context.Context, opus int32) ([]repository.AccountPastWork, error)

func (fq fakeQuerier) CreateLikes(ctx context.Context, arg repository.CreateLikesParams) (repository.Like, error)
func (fq fakeQuerier) DeleteLikesByID(ctx context.Context, arg repository.DeleteLikesByIDParams) (repository.Like, error)
func (fq fakeQuerier) GetLikeStatusByID(ctx context.Context, arg repository.GetLikeStatusByIDParams) (repository.Like, error)
func (fq fakeQuerier) ListLikesByID(ctx context.Context, accountID string) ([]repository.Like, error)
func (fq fakeQuerier) GetListCountByOpus(ctx context.Context, opus int32) (int64, error)

func (fq fakeQuerier) CreatePastWorkFrameworks(ctx context.Context, arg repository.CreatePastWorkFrameworksParams) (repository.PastWorkFramework, error)
func (fq fakeQuerier) DeletePastWorkFrameworksByOpus(ctx context.Context, opus int32) error

func (fq fakeQuerier) CreatePastWorkTags(ctx context.Context, arg repository.CreatePastWorkTagsParams) (repository.PastWorkTag, error)
func (fq fakeQuerier) DeletePastWorkTagsByOpus(ctx context.Context, opus int32) error
func (fq fakeQuerier) ListPastWorkTagsByOpus(ctx context.Context, opus int32) ([]repository.PastWorkTag, error)

func (fq fakeQuerier) GetRoomsAccountsByID(ctx context.Context, roomID string) ([]repository.GetRoomsAccountsByIDRow, error)
func (fq fakeQuerier) CreateRoomsAccounts(ctx context.Context, arg repository.CreateRoomsAccountsParams) (repository.RoomsAccount, error)
func (fq fakeQuerier) DeleteRoomsAccountsByID(ctx context.Context, arg repository.DeleteRoomsAccountsByIDParams) error
