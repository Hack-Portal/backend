package casbin

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	jsonadapter "github.com/casbin/json-adapter/v2"
	"github.com/labstack/echo/v4"
)

func init() {
	// ここでconfを読み込む

}

var casbinModel string = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act, eft

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`

var casbinPolicy []byte = []byte(`[{"PType":"p","V0":"guest","V1":"/v1/hackathons","V2":"GET", "V3":"allow"},{"PType":"p","V0":"guest","V1":"/v1/status_tags","V2":"GET", "V3":"allow"},{"PType":"p","V0":"hack-portal-operator","V1":"/v1/hackathons","V2":"GET", "V3":"allow"},{"PType":"p","V0":"hack-portal-operator","V1":"/v1/hackathons","V2":"POST", "V3":"allow"},{"PType":"p","V0":"hack-portal-operator","V1":"/v1/hackathons/*","V2":"PUT", "V3":"allow"},{"PType":"p","V0":"hack-portal-operator","V1":"/v1/status_tags","V2":"GET", "V3":"allow"},{"PType":"p","V0":"hack-portal-operator","V1":"/v1/status_tags","V2":"POST", "V3":"allow"},{"PType":"p","V0":"hack-portal-operator","V1":"/v1/status_tags/*","V2":"PUT", "V3":"allow"},{"PType":"p","V0":"admin","V1":"*","V2":"*", "V3":"allow"}]`)

const (
	ErrDeny = "deny"
)

func Authorization() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		m, err := model.NewModelFromString(casbinModel)
		if err != nil {
			log.Fatal("model from string error :", err)
		}

		a := jsonadapter.NewAdapter(&casbinPolicy)
		e, err := casbin.NewEnforcer(m, a)
		if err != nil {
			log.Fatal("enforver error :", err)
		}

		return func(c echo.Context) error {
			userRoles := []string{
				"admin",
				//"hack-portal-operator",
				// "guest",
			}

			authorized := false
			obj := c.Request().URL.Path
			act := c.Request().Method

			for _, sub := range userRoles {
				allowed, reason, err := e.EnforceEx(sub, obj, act)
				log.Println(allowed, reason, err)
				if err != nil {
					log.Fatal(err)
				}

				if allowed {
					authorized = true
					continue
				}

				if len(reason) == 0 {
					log.Println(reason)
					continue
				}

				if reason[3] == ErrDeny {
					authorized = false
					break
				}
			}

			if !authorized {
				return echo.ErrForbidden
			}

			return next(c)
		}
	}
}
