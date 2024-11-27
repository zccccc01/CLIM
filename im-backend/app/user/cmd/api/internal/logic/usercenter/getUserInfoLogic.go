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

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get user info
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userID, err := ctxdata.GetUidFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 检查 userID 是否有效
	if userID == nil {
		l.Logger.Error("userID is missing or invalid")
		return nil, errors.New("userID is missing or invalid")
	}

	getUserInfoLogic, err := l.svcCtx.UsercenterRpc.GetUser(l.ctx, &usercenter.GetUserReq{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserInfoResp{
		UserInfo: types.UserProfile{
			UserID:       getUserInfoLogic.UserPro.UserID,
			AvatarURL:    getUserInfoLogic.UserPro.AvatarURL,
			OnlineStatus: getUserInfoLogic.UserPro.OnlineStatus,
			Bio:          getUserInfoLogic.UserPro.Bio,
			Birthday:     getUserInfoLogic.UserPro.Birthday,
			Gender:       getUserInfoLogic.UserPro.Gender,
			Location:     getUserInfoLogic.UserPro.Location,
			LastSeenAt:   getUserInfoLogic.UserPro.LastSeenAt,
		},
	}

	return resp, nil
}
