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

type GetUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserProfileLogic) GetUserProfile(in *pb.GetUserProfileReq) (*pb.GetUserProfileResp, error) {
	// 检查 userID 是否有效
	if in.UserID == nil {
		l.Logger.Error("userID is missing or invalid")
		return nil, errors.New("userID is missing or invalid")
	}

	// 1. 查询用户配置信息
	var userProfile model.UserProfile
	result := l.svcCtx.DB.Where("user_id = ?", in.UserID).First(&userProfile)
	if result.Error != nil {
		l.Logger.Debug("failure to get user profile, err: ", result.Error)
		return nil, result.Error
	}

	return &usercenter.GetUserProfileResp{
		UserProfile: &pb.UserProfile{
			UserID:       userProfile.UserID,
			AvatarURL:    userProfile.AvatarURL,
			OnlineStatus: userProfile.OnlineStatus,
			Bio:          userProfile.Bio,
			Birthday:     userProfile.Birthday.String(),
			Gender:       userProfile.Gender,
			Location:     userProfile.Location,
			LastSeenAt:   userProfile.LastSeenAt.String(),
		},
	}, nil
}
