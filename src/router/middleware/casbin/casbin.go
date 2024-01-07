package casbin

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/Hack-Portal/backend/src/router/middleware/auth"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	jsonadapter "github.com/casbin/json-adapter/v2"
	"github.com/labstack/echo/v4"
)

var casbinModel string

const (
	ErrDeny = "deny"
)

func init() {
	// ここでconfを読み込む
	f, err := os.Open("casbin_model.conf")
	if err != nil {
		log.Fatal("casbin model file open error :", err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Fatal("casbin model file stat error :", err)
	}

	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	if err != nil {
		log.Fatal("casbin model file read error :", err)
	}

	casbinModel = string(buf)
}

type RBACPolicy interface {
	RBACPermission() echo.MiddlewareFunc
}

type RBAC struct {
	rbacRepo dai.RBACPolicyDai
}

func NewRBAC(repo dai.RBACPolicyDai) RBACPolicy {
	return &RBAC{
		rbacRepo: repo,
	}
}

func (rbac *RBAC) RBACPermission() echo.MiddlewareFunc {
	return rbac.rbacPermission
}

func (rbac *RBAC) rbacPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		roleID, ok := c.Get(auth.RequestRoleID).(int)
		if !ok {
			return echo.ErrInternalServerError
		}
		policy, err := rbac.rbacRepo.FindRoleByRole(c.Request().Context(), roleID)
		if err != nil {
			return echo.ErrInternalServerError
		}

		data, err := json.Marshal(policy)
		if err != nil {
			return echo.ErrInternalServerError
		}

		m, err := model.NewModelFromString(casbinModel)
		if err != nil {
			return echo.ErrInternalServerError
		}
		a := jsonadapter.NewAdapter(&data)

		e, err := casbin.NewEnforcer(m, a)
		if err != nil {
			return echo.ErrInternalServerError
		}

		if !rbac.checkPolicy(e, strconv.Itoa(roleID), c.Request().URL.Path, c.Request().Method) {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func (rbac *RBAC) checkPolicy(e *casbin.Enforcer, roleID, obj, act string) bool {
	allowed, reason, err := e.EnforceEx(roleID, obj, act)
	if err != nil {
		return false
	}

	if len(reason) == 0 || reason[3] == ErrDeny {
		return false
	}

	return allowed
}
