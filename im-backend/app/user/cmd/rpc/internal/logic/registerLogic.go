package logic

import (
	"context"
	"errors"
	"fmt"

	"CLIM/app/user/cmd/rpc/internal/svc"
	"CLIM/app/user/cmd/rpc/pb"
	"CLIM/app/user/cmd/rpc/usercenter"
	"CLIM/app/user/dao/model"

	"github.com/gofrs/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 1. 判断手机号和邮箱是否已存在
	var user model.User
	result := l.svcCtx.DB.Where("phone = ? and email = ?", in.Phone, in.Email).First(&user)
	if result.Error == nil {
		if user.Phone == in.Phone {
			l.Logger.Info("phone already exists", logx.Field("phone", in.Phone))
			return nil, fmt.Errorf("phone already in use")
		}
		if user.Email == in.Email {
			l.Logger.Info("email already exists", logx.Field("email", in.Email))
			return nil, fmt.Errorf("email already in use")
		}
	}
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		l.Logger.Error("failed to check phone or email", logx.Field("error", result.Error))
		return nil, fmt.Errorf("failed to check phone or email")
	}
	// 2. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), 14)
	if err != nil {
		l.Logger.Error("failed to hash password", logx.Field("error", err))
		return nil, fmt.Errorf("failed to hash password")
	}
	// 3. 生成基于时间戳的有序 UUID (UUIDv1)
	userID, err := uuid.NewV1()
	if err != nil {
		l.Logger.Error("UUID generation failed", logx.Field("error", err))
		return nil, fmt.Errorf("failed to generate UUID")
	}
	// 4. 创建用户实例
	user = model.User{
		UserID:   userID.Bytes(), // 将 UUID 转换为字节切片存储
		Username: in.UserName,
		Password: string(hashedPassword), // 存储哈希后的密码
		Email:    in.Email,
		Phone:    in.Phone,
	}
	// 5. 插入用户数据到数据库
	if err := l.svcCtx.DB.Create(&user).Error; err != nil {
		l.Logger.Error("Failed to insert user", logx.Field("error", err))
		return nil, fmt.Errorf("failed to register user")
	}
	// 6. 生成令牌,这样服务就不会在内部调用rpc
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserID: user.UserID,
	})
	if err != nil {
		l.Logger.Error("Failed to generate token", logx.Field("error", err))
		return nil, fmt.Errorf("failed to generate token")
	}

	return &usercenter.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
