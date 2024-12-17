package usercenter

import (
	"context"
	"errors"

	"CLIM/app/user/cmd/api/internal/svc"
	"CLIM/app/user/cmd/api/internal/types"
	"CLIM/app/user/cmd/rpc/pb"
	"CLIM/app/user/cmd/rpc/usercenter"
	"CLIM/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// set userProfile
func NewSetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserProfileLogic {
	return &SetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserProfileLogic) SetUserProfile(req *types.SetUserProfileReq) (resp *types.SetUserProfileResp, err error) {
	userID, err := ctxdata.GetUidFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 检查 userID 是否有效
	if userID == nil {
		l.Logger.Error("userID is missing or invalid")
		return nil, errors.New("userID is missing or invalid")
	}

	_, err = l.svcCtx.UsercenterRpc.SetUserProfile(l.ctx, &usercenter.SetUserProfileReq{
		UserProfile: &pb.UserProfile{
			UserID:       userID,
			AvatarURL:    req.UserProfile.AvatarURL,
			OnlineStatus: req.UserProfile.OnlineStatus,
			Bio:          req.UserProfile.Bio,
			Birthday:     req.UserProfile.Birthday,
			Gender:       req.UserProfile.Gender,
			Location:     req.UserProfile.Location,
			LastSeenAt:   req.UserProfile.LastSeenAt,
		},
	})
	if err != nil {
		return nil, err
	}

	return
}
