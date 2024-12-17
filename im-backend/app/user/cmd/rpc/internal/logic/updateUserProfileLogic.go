package logic

import (
	"context"
	"time"

	"CLIM/app/user/cmd/rpc/internal/svc"
	"CLIM/app/user/cmd/rpc/pb"
	"CLIM/app/user/dao/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(in *pb.UpdateUserProfileReq) (*pb.UpdateUserProfileResp, error) {
	// 1. 查询用户信息
	var userProfile model.UserProfile
	result := l.svcCtx.DB.Where("user_id = ?", in.UserProfile.UserID).First(&userProfile)
	if result.Error != nil {
		l.Logger.Debug("user not found")
		return nil, result.Error
	}
	// 2. 用map存储要更新的字段
	values := map[string]interface{}{
		"AvatarURL":    in.UserProfile.AvatarURL,
		"OnlineStatus": in.UserProfile.OnlineStatus,
		"Bio":          in.UserProfile.Bio,
		"Birthday":     in.UserProfile.Birthday,
		"Gender":       in.UserProfile.Gender,
		"Location":     in.UserProfile.Location,
		"LastSeenAt":   time.Now(),
	}
	// 3. 更新数据库
	result = l.svcCtx.DB.Model(&userProfile).Updates(values)
	if result.Error != nil {
		l.Logger.Debug("update user profile failed", result.Error)
		return nil, result.Error
	}

	return &pb.UpdateUserProfileResp{}, nil
}
