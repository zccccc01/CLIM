package usercenter

import (
	"context"

	"CLIM/app/user/cmd/api/internal/svc"
	"CLIM/app/user/cmd/api/internal/types"
	"CLIM/app/user/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginLogic, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResp{
		AccessToken:  loginLogic.AccessToken,
		AccessExpire: loginLogic.AccessExpire,
		RefreshAfter: loginLogic.RefreshAfter,
	}

	return resp, nil
}
