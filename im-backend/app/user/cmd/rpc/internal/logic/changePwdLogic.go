package logic

import (
	"context"
	"fmt"

	"CLIM/app/user/cmd/rpc/internal/svc"
	"CLIM/app/user/cmd/rpc/pb"
	"CLIM/app/user/dao/model"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type ChangePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePwdLogic {
	return &ChangePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePwdLogic) ChangePwd(in *pb.ChangePwdReq) (*pb.ChangePwdResp, error) {
	// 1. 查询用户是否存在
	var user model.User
	result := l.svcCtx.DB.Where("phone = ?", in.Phone).First(&user)
	if result.Error != nil {
		l.Logger.Info("user not exist")
		return nil, result.Error
	}
	// 2. 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword)); err != nil {
		l.Logger.Info("password not correct")
		return nil, fmt.Errorf("password error")
	}
	// 3. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Error("failed to hash password", logx.Field("error", err))
		return nil, fmt.Errorf("failed to hash password")
	}
	// 4. 更新密码
	result = l.svcCtx.DB.Model(&user).Update("password", hashedPassword)
	if result.Error != nil {
		l.Logger.Info("update password failed")
		return nil, result.Error
	}

	return &pb.ChangePwdResp{}, nil
}
