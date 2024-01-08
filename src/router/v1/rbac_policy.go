package v1

import (
	"github.com/Hack-Portal/backend/src/adapters/controllers"
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/adapters/presenters"
	"github.com/Hack-Portal/backend/src/usecases/interactors"
)

func (r *v1router) rbacPolicy() {
	rbac := r.v1.Group("/rbac")

	rp := controllers.NewRbacPolicyController(
		interactors.NewRbacPolicyInteractor(
			gateways.NewRbacPolicyGateway(r.db),
			gateways.NewRoleGateway(r.db),
			presenters.NewRbacPolicyPresenter(),
		),
	)

	// 作成
	rbac.POST("", rp.Create)
	// 全取得
	rbac.GET("", rp.ReadAll)
	// 全削除
	rbac.DELETE("", rp.Delete)
	// 個別削除
	rbac.DELETE("/:policy_id", rp.DeleteAll)
}
