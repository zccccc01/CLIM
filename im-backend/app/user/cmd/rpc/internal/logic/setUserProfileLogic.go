package logic

import (
	"context"
	"fmt"
	"time"

	"CLIM/app/user/cmd/rpc/internal/svc"
	"CLIM/app/user/cmd/rpc/pb"
	"CLIM/app/user/dao/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserProfileLogic {
	return &SetUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserProfileLogic) SetUserProfile(in *pb.SetUserProfileReq) (*pb.SetUserProfileResp, error) {
	// 1. 将用户信息存入数据库
	var userProfile model.UserProfile
	parseTime, err := time.Parse("2006-01-02", in.UserProfile.Birthday)
	if err != nil {
		l.Logger.Debug(err)
		return nil, fmt.Errorf("parse birthday failed")
	}
	// 2. 设置上海时区
	shanghaiLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		l.Logger.Debug(err)
		return nil, fmt.Errorf("load shanghai location failed")
	}

	parseTime = time.Date(parseTime.Year(), parseTime.Month(), parseTime.Day(), 0, 0, 0, 0, shanghaiLocation)

	userProfile = model.UserProfile{
		UserID:       in.UserProfile.UserID,
		AvatarURL:    in.UserProfile.AvatarURL,
		OnlineStatus: in.UserProfile.OnlineStatus,
		Bio:          in.UserProfile.Bio,
		Birthday:     parseTime,
		Gender:       in.UserProfile.Gender,
		Location:     in.UserProfile.Location,
		LastSeenAt:   time.Now(),
	}

	if err = l.svcCtx.DB.Create(&userProfile).Error; err != nil {
		l.Logger.Info(err)
		return nil, fmt.Errorf("create user profile failed")
	}
	return &pb.SetUserProfileResp{}, nil
}
