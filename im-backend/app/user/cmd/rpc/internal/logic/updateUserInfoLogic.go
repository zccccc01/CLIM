package logic

import (
	"context"

	"CLIM/app/user/cmd/rpc/internal/svc"
	"CLIM/app/user/cmd/rpc/pb"
	"CLIM/app/user/dao/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	// 1. 查询用户信息
	var user model.User
	result := l.svcCtx.DB.Where("user_id = ?", in.UserID).First(&user)
	if result.Error != nil {
		l.Logger.Debug("user not found")
		return nil, result.Error
	}
	// 2. 用map存储要更新的字段
	values := map[string]interface{}{
		"UserName": in.UserName,
		"Email":    in.Email,
		"Phone":    in.Phone,
	}
	// 3. 更新数据库
	result = l.svcCtx.DB.Model(&user).Where("user_id = ?", in.UserID).Updates(values)
	if result.Error != nil {
		l.Logger.Debug("update user info failed")
		return nil, result.Error
	}

	return &pb.UpdateUserInfoResp{}, nil
}
