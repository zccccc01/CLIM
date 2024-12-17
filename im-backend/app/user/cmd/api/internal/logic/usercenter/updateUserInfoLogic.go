package usercenter

import (
	"context"
	"errors"

	"CLIM/app/user/cmd/api/internal/svc"
	"CLIM/app/user/cmd/api/internal/types"
	"CLIM/app/user/cmd/rpc/usercenter"
	"CLIM/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update user info
func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoResp, err error) {
	userID, err := ctxdata.GetUidFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 检查 userID 是否有效
	if userID == nil {
		l.Logger.Error("userID is missing or invalid")
		return nil, errors.New("userID is missing or invalid")
	}

	_, err = l.svcCtx.UsercenterRpc.UpdateUserInfo(l.ctx, &usercenter.UpdateUserInfoReq{
		UserID:   userID,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return
}
