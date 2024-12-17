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

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get userProfile
func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileLogic) GetUserProfile(req *types.UserProfileReq) (resp *types.UserProfileResp, err error) {
	userID, err := ctxdata.GetUidFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 检查 userID 是否有效
	if userID == nil {
		l.Logger.Error("userID is missing or invalid")
		return nil, errors.New("userID is missing or invalid")
	}

	getUserProfileLogic, err := l.svcCtx.UsercenterRpc.GetUserProfile(l.ctx, &usercenter.GetUserProfileReq{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserProfileResp{
		UserProfile: types.UserProfile{
			AvatarURL:    getUserProfileLogic.UserProfile.AvatarURL,
			OnlineStatus: getUserProfileLogic.UserProfile.OnlineStatus,
			Bio:          getUserProfileLogic.UserProfile.Bio,
			Birthday:     getUserProfileLogic.UserProfile.Birthday,
			Gender:       getUserProfileLogic.UserProfile.Gender,
			Location:     getUserProfileLogic.UserProfile.Location,
			LastSeenAt:   getUserProfileLogic.UserProfile.LastSeenAt,
		},
	}

	return resp, nil
}
