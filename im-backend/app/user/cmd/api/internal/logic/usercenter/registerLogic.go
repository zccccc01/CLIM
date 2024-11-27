package usercenter

import (
	"context"

	"CLIM/app/user/cmd/api/internal/svc"
	"CLIM/app/user/cmd/api/internal/types"
	"CLIM/app/user/cmd/rpc/usercenter"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerResp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.RegisterResp{
		AccessToken:  registerResp.AccessToken,
		AccessExpire: registerResp.AccessExpire,
		RefreshAfter: registerResp.RefreshAfter,
	}

	return resp, nil
}
