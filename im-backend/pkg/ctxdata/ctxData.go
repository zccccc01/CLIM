package ctxdata

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"

	"github.com/zeromicro/go-zero/core/logx"
)

var CtxKeyJwtUserID = "user_id"

func GetUidFromCtx(ctx context.Context) ([]byte, error) {
	// 1. 从上下文中检索用户ID
	userIDBase64, ok := ctx.Value(CtxKeyJwtUserID).(string)
	if !ok || userIDBase64 == "" {
		logx.Errorf("missing or invalid user ID in context. Context type: %v, Context value: %+v", reflect.TypeOf(ctx), ctx)
		return nil, errors.New("missing or invalid user ID in context")
	}

	// 记录检索到的userIDBase64以便于调试
	logx.Infof("retrieved userIDBase64 from context: %s", userIDBase64)

	// 2. 将Base64编码的用户ID解码为二进制
	userID, err := base64.StdEncoding.DecodeString(userIDBase64)
	if err != nil {
		logx.Errorf("failed to decode user ID from Base64: %v, input: %s", err, userIDBase64)
		return nil, err
	}

	logx.Infof("successfully decoded userID to binary: %v", fmt.Sprintf("%x", userID))

	return userID, nil
}
