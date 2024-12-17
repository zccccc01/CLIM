package usercenter

import (
	"context"

	"CLIM/app/user/cmd/api/internal/svc"
	"CLIM/app/user/cmd/api/internal/types"
	"CLIM/app/user/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// change pwd
func NewChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePwdLogic {
	return &ChangePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePwdLogic) ChangePwd(req *types.ChangePwdReq) (resp *types.ChangePwdResp, err error) {
	_, err = l.svcCtx.UsercenterRpc.ChangePwd(l.ctx, &usercenter.ChangePwdReq{
		Phone:       req.Phone,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}

	return
}
