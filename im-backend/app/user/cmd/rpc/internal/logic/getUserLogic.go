package logic

import (
	"context"

	"CLIM/app/user/cmd/rpc/internal/svc"
	"CLIM/app/user/cmd/rpc/pb"
	"CLIM/app/user/cmd/rpc/usercenter"
	"CLIM/app/user/dao/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.GetUserReq) (*pb.GetUserResp, error) {
	// 检查 userID 是否有效
	if in.UserID == nil {
		l.Logger.Debug("userID is missing or invalid")
		return nil, errors.New("userID is missing or invalid")
	}

	// 1. 查询用户信息
	var user model.User
	result := l.svcCtx.DB.Where("user_id = ?", in.UserID).First(&user)
	if result.Error != nil {
		l.Logger.Debug("failure to get user info, err: ", result.Error)
		return nil, result.Error
	}

	return &usercenter.GetUserResp{
		User: &pb.User{
			UserID:   user.UserID,
			UserName: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}, nil
}
