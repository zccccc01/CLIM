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

type UpdateUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update userProfile
func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(req *types.UpdateUserProfileReq) (resp *types.UpdateUserProfileResp, err error) {
	userID, err := ctxdata.GetUidFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 检查 userID 是否有效
	if userID == nil {
		l.Logger.Error("userID is missing or invalid")
		return nil, errors.New("userID is missing or invalid")
	}

	_, err = l.svcCtx.UsercenterRpc.UpdateUserProfile(l.ctx, &usercenter.UpdateUserProfileReq{
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
