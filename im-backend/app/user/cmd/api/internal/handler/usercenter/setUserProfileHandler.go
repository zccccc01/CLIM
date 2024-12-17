package usercenter

import (
	"net/http"

	"CLIM/app/user/cmd/api/internal/logic/usercenter"
	"CLIM/app/user/cmd/api/internal/svc"
	"CLIM/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// set userProfile
func SetUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUserProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := usercenter.NewSetUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.SetUserProfile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
