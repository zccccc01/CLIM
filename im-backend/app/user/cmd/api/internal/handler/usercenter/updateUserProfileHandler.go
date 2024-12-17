package usercenter

import (
	"net/http"

	"CLIM/app/user/cmd/api/internal/logic/usercenter"
	"CLIM/app/user/cmd/api/internal/svc"
	"CLIM/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// update userProfile
func UpdateUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := usercenter.NewUpdateUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserProfile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
