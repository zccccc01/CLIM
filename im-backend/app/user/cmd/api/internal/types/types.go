// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type ChangePwdReq struct {
	Phone       string `json:"phone"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePwdResp struct {
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type RegisterReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type RegisterResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type SetUserProfileReq struct {
	UserProfile UserProfile `json:"user_profile"`
}

type SetUserProfileResp struct {
}

type UpdateUserInfoReq struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UpdateUserInfoResp struct {
}

type UpdateUserProfileReq struct {
	UserProfile UserProfile `json:"user_profile"`
}

type UpdateUserProfileResp struct {
}

type User struct {
	UserID   []byte `json:"user_id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"user_info"`
}

type UserProfile struct {
	AvatarURL    string `json:"avatar_url"`
	OnlineStatus int32  `json:"online_status"` // 0: offline, 1: online
	Bio          string `json:"bio"`           // Short biography or status message
	Birthday     string `json:"birthday"`      // Users date of birth
	Gender       string `json:"gender"`        // User gender
	Location     string `json:"location"`      // Users location or timezone
	LastSeenAt   string `json:"last_seen_at"`  // Last seen timestamp for the user
}

type UserProfileReq struct {
}

type UserProfileResp struct {
	UserProfile UserProfile `json:"user_profile"`
}
