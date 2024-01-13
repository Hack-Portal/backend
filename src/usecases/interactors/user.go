package interactors

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/Hack-Portal/backend/src/utils/password"
	"github.com/Hack-Portal/backend/src/utils/random"
	"github.com/labstack/echo/v4"
)

type userInteractor struct {
	userRepo dai.UsersDai
	roleRepo dai.RoleDai
	output   ports.UserOutputBoundary
}

// NewUserInteractor はUserに関するユースケースを生成します
func NewUserInteractor(userRepo dai.UsersDai, roleRepo dai.RoleDai, output ports.UserOutputBoundary) ports.UserInputBoundary {
	return &userInteractor{
		userRepo: userRepo,
		roleRepo: roleRepo,
		output:   output,
	}
}

// InitAdmin は管理者を初期化します
func (u *userInteractor) InitAdmin(ctx context.Context, in request.InitAdmin) (int, *response.User) {
	// Tokenの検証 => 失敗したらTokenを変更して返す TODO:スケールしない構成になっているから、Redisに保存するようにする
	if in.InitAdminToken != config.Config.Server.AdminInitPassword {
		config.Config.Server.AdminInitPassword = random.AlphaNumeric(30)
		log.Printf("AdminInitPassword changed to [%s] \n", config.Config.Server.AdminInitPassword)
		return u.output.PresentInitAdmin(ctx, &ports.OutputInitAdminData{
			Error:    echo.ErrBadRequest,
			Response: nil,
		})
	}

	pass := random.AlphaNumeric(30)
	hashed, err := password.HashPassword(pass)
	if err != nil {
		return u.output.PresentInitAdmin(ctx, &ports.OutputInitAdminData{
			Error:    err,
			Response: nil,
		})
	}

	arg := &models.User{
		UserID:   random.AlphaNumeric(30),
		Name:     in.Name,
		Password: hashed,
		Role:     models.RoleAdmin,
	}

	_, err = u.userRepo.Create(ctx, arg)
	if err != nil {
		return u.output.PresentInitAdmin(ctx, &ports.OutputInitAdminData{
			Error:    err,
			Response: nil,
		})
	}

	return u.output.PresentInitAdmin(ctx, &ports.OutputInitAdminData{
		Error:    nil,
		Response: &response.User{UserID: arg.UserID, Name: arg.Name, Password: pass},
	})
}

// Login はログイン用だが、現状はユーザーの情報とロールを返すだけ
func (u *userInteractor) Login(ctx context.Context, in request.Login) (int, *response.Login) {
	user, err := u.userRepo.FindByID(ctx, in.UserID)
	if err != nil {
		return u.output.PresentLogin(ctx, ports.NewOutput[*response.Login](err, nil))
	}

	if err := password.CheckPassword(in.Password, user.Password); err != nil {
		log.Println("password check error", err)
		return u.output.PresentLogin(ctx, ports.NewOutput[*response.Login](err, nil))
	}

	role, err := u.roleRepo.FindByID(ctx, int64(user.Role))
	if err != nil {
		log.Println("find role error", err)
		return u.output.PresentLogin(ctx, ports.NewOutput[*response.Login](err, nil))
	}

	token := base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user.UserID, in.Password)))

	return u.output.PresentLogin(ctx, ports.NewOutput[*response.Login](nil, &response.Login{
		UserID: user.UserID,
		Name:   user.Name,
		Role:   role.Role,
		Token:  string(token),
	}))
}
