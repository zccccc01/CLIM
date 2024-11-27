package logic

import (
	"context"
	"fmt"

	"CLIM/app/user/cmd/rpc/internal/svc"
	"CLIM/app/user/cmd/rpc/pb"
	"CLIM/app/user/cmd/rpc/usercenter"
	"CLIM/app/user/dao/model"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 1.按手机号查询用户
	var user model.User
	result := l.svcCtx.DB.Where("phone=?", in.Phone).First(&user)
	if result.Error != nil {
		l.Logger.Info("user not exist", logx.Field("error", result.Error))
		return nil, fmt.Errorf("user not exist")
	}
	// 2.校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		l.Logger.Info("password not match")
		return nil, fmt.Errorf("password error")
	}
	// 3. 生成令牌,这样服务就不会在内部调用rpc
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserID: user.UserID,
	})
	if err != nil {
		l.Logger.Error("Failed to generate token", logx.Field("error", err))
		return nil, fmt.Errorf("failed to generate token")
	}

	return &usercenter.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
